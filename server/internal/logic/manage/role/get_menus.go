package role

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role_menu"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/role"
)

type GetMenus struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetMenus(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetMenus {
	return &GetMenus{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *GetMenus) GetMenus(req *types.GetMenusRequest) (resp *types.GetMenusResponse, err error) {
	menus, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.Condition{
		Field:    manage_role_menu.RoleUuid,
		Operator: condition.Equal,
		Value:    req.RoleUuid,
	})
	if err != nil {
		return
	}

	var menuUuids []string
	for _, menu := range menus {
		menuUuids = append(menuUuids, menu.MenuUuid)
	}
	resp = &types.GetMenusResponse{MenuUuids: menuUuids}

	return
}
