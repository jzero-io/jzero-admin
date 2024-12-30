package route

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/jzero-io/jzero-contrib/nullx"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/auth"
	"github.com/jzero-io/jzero-admin/server/internal/logic/manage/menu"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_menu"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/route"
)

type GetUserRoutes struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetUserRoutes(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetUserRoutes {
	return &GetUserRoutes{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *GetUserRoutes) GetUserRoutes(req *types.GetUserRoutesRequest) (resp *types.GetUserRoutesResponse, err error) {
	resp = &types.GetUserRoutesResponse{
		Routes: []types.Route{},
	}
	info, err := auth.Info(l.ctx)
	if err != nil {
		return nil, err
	}

	if info.RoleIds == nil {
		return resp, nil
	}

	roleMenus, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		In("role_id", info.RoleIds).
		Build()...)
	if err != nil {
		return nil, err
	}

	var menuIds []uint64
	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, uint64(roleMenu.MenuId))
	}
	uniqMenuIds := lo.Uniq(menuIds)

	if len(uniqMenuIds) == 0 {
		return resp, nil
	}

	menus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		In("id", uniqMenuIds).
		NotEqual("menu_type", "3").
		Build()...)
	if err != nil {
		return nil, err
	}
	list := buildRouteTree(convert(menus), 0)

	resp.Routes = list

	// get home
	for _, rm := range roleMenus {
		if cast.ToBool(rm.IsHome) {
			for _, m := range menus {
				if m.Id == uint64(rm.MenuId) {
					resp.Home = m.RouteName
				}
			}
		}
	}

	return
}

func convert(list []*manage_menu.ManageMenu) []*types.Route {
	var records []*types.Route
	for _, item := range list {
		var route types.Route
		var query []types.Query

		menu.Unmarshal(item.Query.String, &query)

		route = types.Route{
			Id:       int64(item.Id),
			ParentId: item.ParentId,
			Name:     item.RouteName,
			Path:     item.RoutePath,
			Meta: types.RouteMeta{
				Title:           item.RouteName,
				I18nKey:         item.I18nKey,
				Icon:            item.Icon,
				Order:           int(item.Order),
				HideInMenu:      cast.ToBool(item.HideInMenu),
				ActiveMenu:      item.ActiveMenu.String,
				MultiTab:        cast.ToBool(item.MultiTab),
				FixedIndexInTab: nullx.NewInt(item.FixedIndexInTab).Ptr(),
				KeepAlive:       cast.ToBool(item.KeepAlive),
				Constant:        cast.ToBool(item.Constant),
				Href:            item.Href.String,
				Query:           query,
			},
			Component: item.Component,
			Redirect:  "",
		}
		if item.Component == "view.iframe-page" {
			route.Props = map[string]any{
				"url": item.Href.String,
			}
			route.Meta.Href = ""
		}
		if item.IconType == "2" {
			route.Meta.LocalIcon = item.Icon
			route.Meta.Icon = ""
		}
		records = append(records, &route)
	}
	return records
}

func buildRouteTree(routes []*types.Route, parentId int64) []types.Route {
	var result []types.Route
	for _, route := range routes {
		if route.ParentId == parentId {
			subRoute := *route
			subRoute.Children = buildRouteTree(routes, route.Id)
			result = append(result, subRoute)
		}
	}
	return result
}
