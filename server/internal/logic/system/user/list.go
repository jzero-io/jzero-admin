package user

import (
	"context"
	"time"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/jzero-io/jzero-contrib/nullx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"

	"server/internal/svc"
	types "server/internal/types/system/user"
)

type List struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewList(ctx context.Context, svcCtx *svc.ServiceContext) *List {
	return &List{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *List) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	users, total, err := l.svcCtx.Model.SystemUser.PageByCondition(l.ctx, nil, condition.Condition{
		Operator: condition.Limit,
		Value:    req.Size,
	}, condition.Condition{
		Operator: condition.Offset,
		Value:    (req.Current - 1) * req.Size,
	}, condition.Condition{
		Skip:     req.Username == "",
		Field:    "username",
		Operator: condition.Like,
		Value:    "%" + req.Username + "%",
	}, condition.Condition{
		Skip:     req.UserGender == "",
		Field:    "gender",
		Operator: condition.Equal,
		Value:    req.UserGender,
	}, condition.Condition{
		Skip:     req.NickName == "",
		Field:    "nickname",
		Operator: condition.Like,
		Value:    "%" + req.NickName + "%",
	}, condition.Condition{
		Skip:     req.UserPhone == "",
		Field:    "phone",
		Operator: condition.Like,
		Value:    "%" + req.UserPhone + "%",
	}, condition.Condition{
		Skip:     req.UserEmail == "",
		Field:    "email",
		Operator: condition.Like,
		Value:    "%" + req.UserEmail + "%",
	}, condition.Condition{
		Skip:     req.Status == "",
		Field:    "status",
		Operator: condition.Equal,
		Value:    req.Status,
	})

	var records []types.SystemUser
	for _, user := range users {
		records = append(records, types.SystemUser{
			Id:         user.Id,
			Username:   user.Username,
			UserGender: user.Gender,
			NickName:   user.Nickname,
			UserPhone:  nullx.NewString(user.Phone).ValueOrZero(),
			UserEmail:  nullx.NewString(user.Email).ValueOrZero(),
			Status:     user.Status,
			CreateTime: user.CreateTime.Format(time.DateTime),
			UpdateTime: user.UpdateTime.Format(time.DateTime),
		})
	}

	err = mr.MapReduceVoid(func(source chan<- int) {
		for index := range records {
			source <- index
		}
	}, func(item int, writer mr.Writer[types.SystemUser], cancel func(error)) {
		userRoles, err := l.svcCtx.Model.SystemUserRole.FindByCondition(l.ctx, nil, condition.Condition{
			Field:    "user_id",
			Operator: condition.Equal,
			Value:    records[item].Id,
		})
		if err != nil {
			cancel(err)
			return
		}
		var roleIds []int
		for _, userRole := range userRoles {
			roleIds = append(roleIds, int(userRole.RoleId))
		}
		if len(roleIds) == 0 {
			return
		}

		roles, err := l.svcCtx.Model.SystemRole.FindByCondition(l.ctx, nil, condition.Condition{
			Field:    "id",
			Operator: condition.In,
			Value:    roleIds,
		})
		if err != nil {
			cancel(err)
			return
		}
		var roleCodes []string
		for _, role := range roles {
			roleCodes = append(roleCodes, role.Code)
		}
		records[item].UserRoles = roleCodes
	}, func(pipe <-chan types.SystemUser, cancel func(error)) {})

	resp = &types.ListResponse{
		Records: records,
		PageResponse: types.PageResponse{
			Current: req.Current,
			Size:    req.Size,
			Total:   total,
		},
	}
	return
}
