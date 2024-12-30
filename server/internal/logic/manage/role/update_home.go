package role

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/role"
)

type UpdateHome struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUpdateHome(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UpdateHome {
	return &UpdateHome{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *UpdateHome) UpdateHome(req *types.UpdateHomeRequest) (resp *types.Empty, err error) {
	menu, err := l.svcCtx.Model.ManageMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
		Equal("route_name", req.Home).
		Build()...)
	if err != nil {
		return nil, err
	}

	// 找到旧 home
	oldRoleHomeMenu, err := l.svcCtx.Model.ManageRoleMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
		Equal("role_id", req.RoleId).
		Equal("is_home", true).
		Build()...)
	if err == nil {
		oldRoleHomeMenu.IsHome = cast.ToInt64(false)
		err = l.svcCtx.Model.ManageRoleMenu.Update(l.ctx, nil, oldRoleHomeMenu)
		if err != nil {
			return nil, err
		}
	}

	roleMenu, err := l.svcCtx.Model.ManageRoleMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
		Equal("role_id", req.RoleId).
		Equal("menu_id", menu.Id).
		Build()...)
	if err != nil {
		return nil, err
	}

	roleMenu.IsHome = cast.ToInt64(true)

	err = l.svcCtx.Model.ManageRoleMenu.Update(l.ctx, nil, roleMenu)
	if err != nil {
		return nil, err
	}
	return
}
