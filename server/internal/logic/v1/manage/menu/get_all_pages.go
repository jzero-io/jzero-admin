package menu

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_menu"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role_menu"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/menu"
)

type GetAllPages struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetAllPages(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetAllPages {
	return &GetAllPages{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *GetAllPages) GetAllPages(req *types.GetAllPagesRequest) (resp []string, err error) {
	var pages []*manage_menu.ManageMenu
	if req.RoleUuid != "" {
		roleMenus, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.NewChain().
			Equal(manage_role_menu.RoleUuid, req.RoleUuid).
			Build()...)
		if err != nil {
			return nil, err
		}

		var menuUuids []string
		for _, roleMenu := range roleMenus {
			menuUuids = append(menuUuids, roleMenu.MenuUuid)
		}
		uniqMenuUuids := lo.Uniq(menuUuids)

		if len(uniqMenuUuids) == 0 {
			return resp, nil
		}

		pages, err = l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
			In(manage_menu.Uuid, uniqMenuUuids).
			Equal(manage_menu.MenuType, "2").
			Equal(manage_menu.Status, "1").
			NotEqual(manage_menu.HideInMenu, cast.ToInt(true)).
			Build()...)
		if err != nil {
			return nil, err
		}
	} else {
		pages, err = l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
			Equal(manage_menu.MenuType, "2").
			Equal(manage_menu.Status, "1").
			NotEqual(manage_menu.HideInMenu, cast.ToInt(true)).
			Build()...)
		if err != nil {
			return nil, err
		}
	}
	for _, page := range pages {
		resp = append(resp, page.RouteName)
	}
	return resp, nil
}
