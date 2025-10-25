package role

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role_menu"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/role"
)

type GetHome struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetHome(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetHome {
	return &GetHome{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *GetHome) GetHome(req *types.GetHomeRequest) (resp string, err error) {
	roleHomeMenu, err := l.svcCtx.Model.ManageRoleMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
		Equal(manage_role_menu.RoleUuid, req.RoleUuid).
		Equal(manage_role_menu.IsHome, cast.ToInt(true)).
		Build()...)
	if err != nil {
		return "", err
	}
	one, err := l.svcCtx.Model.ManageMenu.FindOneByUuid(l.ctx, nil, roleHomeMenu.MenuUuid)
	if err != nil {
		return "", err
	}
	return one.RouteName, nil
}
