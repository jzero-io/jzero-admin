package user

import (
	"context"
	"time"

	"github.com/guregu/null/v5"
	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/model/system_user"
	"server/internal/model/system_user_role"
	"server/internal/svc"
	types "server/internal/types/system/user"
)

type Add struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdd(ctx context.Context, svcCtx *svc.ServiceContext) *Add {
	return &Add{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Add) Add(req *types.AddRequest) (resp *types.AddResponse, err error) {
	if _, err = l.svcCtx.Model.SystemUser.Insert(l.ctx, &system_user.SystemUser{
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Username:   req.Username,
		Password:   req.Password,
		Nickname:   req.NickName,
		Gender:     req.UserGender,
		Phone:      null.StringFrom(req.UserPhone).NullString,
		Status:     req.Status,
		Email:      null.StringFrom(req.UserEmail).NullString,
	}); err != nil {
		return nil, err
	}
	user, err := l.svcCtx.Model.SystemUser.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}

	var bulk []*system_user_role.SystemUserRole
	var roleIds []uint64
	roles, err := l.svcCtx.Model.SystemRole.FindByCondition(l.ctx, condition.Condition{
		Field:    "code",
		Operator: condition.In,
		Value:    req.UserRoles,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range roles {
		roleIds = append(roleIds, v.Id)
	}

	for _, v := range roleIds {
		bulk = append(bulk, &system_user_role.SystemUserRole{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			UserId:     int64(user.Id),
			RoleId:     int64(v),
		})
	}

	err = l.svcCtx.Model.SystemUserRole.BulkInsert(l.ctx, bulk)

	return
}
