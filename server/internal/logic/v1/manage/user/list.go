package user

import (
	"context"
	"net/http"
	"time"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_user"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_user_role"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/user"
)

type List struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewList(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *List {
	return &List{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *List) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	users, total, err := l.svcCtx.Model.ManageUser.PageByCondition(l.ctx, nil, condition.NewChain().
		Page(req.Current, req.Size).
		OrderByDesc(manage_user.CreateTime).
		Like(manage_user.Username, "%"+req.Username+"%").
		Equal(manage_user.Gender, req.UserGender, condition.WithSkip(req.UserGender == "")).
		Like(manage_user.Nickname, "%"+req.NickName+"%", condition.WithSkip(req.NickName == "")).
		Like(manage_user.Phone, "%"+req.UserPhone+"%", condition.WithSkip(req.UserPhone == "")).
		Like(manage_user.Email, "%"+req.UserEmail+"%", condition.WithSkip(req.UserEmail == "")).
		Equal(manage_user.Status, req.Status, condition.WithSkip(req.Status == "")).
		Build()...)

	var records []types.ManageUser
	for _, user := range users {
		records = append(records, types.ManageUser{
			Uuid:       user.Uuid,
			Username:   user.Username,
			UserGender: user.Gender,
			NickName:   user.Nickname,
			UserPhone:  user.Phone,
			UserEmail:  user.Email,
			Status:     user.Status,
			CreateTime: user.CreateTime.Format(time.DateTime),
			UpdateTime: user.UpdateTime.Format(time.DateTime),
		})
	}

	err = mr.MapReduceVoid(func(source chan<- int) {
		for index := range records {
			source <- index
		}
	}, func(item int, writer mr.Writer[types.ManageUser], cancel func(error)) {
		userRoles, err := l.svcCtx.Model.ManageUserRole.FindByCondition(l.ctx, nil, condition.Condition{
			Field:    manage_user_role.UserUuid,
			Operator: condition.Equal,
			Value:    records[item].Uuid,
		})
		if err != nil {
			cancel(err)
			return
		}
		var roleUuids []string
		for _, userRole := range userRoles {
			roleUuids = append(roleUuids, userRole.RoleUuid)
		}
		if len(roleUuids) == 0 {
			return
		}

		roles, err := l.svcCtx.Model.ManageRole.FindByCondition(l.ctx, nil, condition.Condition{
			Field:    manage_role.Uuid,
			Operator: condition.In,
			Value:    roleUuids,
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
	}, func(pipe <-chan types.ManageUser, cancel func(error)) {})

	if err != nil {
		return nil, err
	}

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
