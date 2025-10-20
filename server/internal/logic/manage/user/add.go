package user

import (
	"context"
	"net/http"
	"time"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_user"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_user_role"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/user"
)

type Add struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAdd(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Add {
	return &Add{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *Add) Add(req *types.AddRequest) (resp *types.AddResponse, err error) {
	if _, err = l.svcCtx.Model.ManageUser.Insert(l.ctx, nil, &manage_user.ManageUser{
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Username:   req.Username,
		Password:   req.Password,
		Nickname:   req.NickName,
		Gender:     req.UserGender,
		Phone:      req.UserPhone,
		Status:     req.Status,
		Email:      req.UserEmail,
	}); err != nil {
		return nil, err
	}
	user, err := l.svcCtx.Model.ManageUser.FindOneByUsername(l.ctx, nil, req.Username)
	if err != nil {
		return nil, err
	}

	var bulk []*manage_user_role.ManageUserRole
	var roleIds []int64
	roles, err := l.svcCtx.Model.ManageRole.FindByCondition(l.ctx, nil, condition.Condition{
		Field:    manage_role.Code,
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
		bulk = append(bulk, &manage_user_role.ManageUserRole{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			UserId:     user.Id,
			RoleId:     v,
		})
	}

	err = l.svcCtx.Model.ManageUserRole.BulkInsert(l.ctx, nil, bulk)

	return
}
