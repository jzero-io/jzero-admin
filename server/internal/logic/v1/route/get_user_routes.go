package route

import (
	"context"
	"net/http"

	"github.com/guregu/null/v5"
	"github.com/jzero-io/jzero-admin/core-engine/helper/auth"
	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/logic/v1/manage/menu"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_menu"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role_menu"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/route"
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

	if info.RoleUuids == nil {
		return resp, nil
	}

	roleMenus, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		In(manage_role_menu.RoleUuid, info.RoleUuids).
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

	menus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		In(manage_menu.Uuid, uniqMenuUuids).
		NotEqual(manage_menu.MenuType, "3").
		Build()...)
	if err != nil {
		return nil, err
	}
	_, uuidMap := convert(menus)
	list := buildRouteTree(menus, uuidMap, "")

	resp.Routes = list

	// get home
	for _, rm := range roleMenus {
		if cast.ToBool(rm.IsHome) {
			for _, m := range menus {
				if m.Uuid == rm.MenuUuid {
					resp.Home = m.RouteName
				}
			}
		}
	}

	return
}

func convert(list []*manage_menu.ManageMenu) ([]*types.Route, map[string]*types.Route) {
	var records []*types.Route
	uuidMap := make(map[string]*types.Route)
	for _, item := range list {
		var route types.Route
		var query []types.Query

		menu.Unmarshal(item.Query, &query)

		route = types.Route{
			Id:       item.Id,
			ParentId: 0,
			Name:     item.RouteName,
			Path:     item.RoutePath,
			Meta: types.RouteMeta{
				Title:           item.RouteName,
				I18nKey:         item.I18nKey,
				Icon:            item.Icon,
				Order:           int(item.Order),
				HideInMenu:      cast.ToBool(item.HideInMenu),
				ActiveMenu:      item.ActiveMenu,
				MultiTab:        cast.ToBool(item.MultiTab),
				FixedIndexInTab: null.NewInt(item.FixedIndexInTab, item.FixedIndexInTab != 0).Ptr(),
				KeepAlive:       cast.ToBool(item.KeepAlive),
				Constant:        cast.ToBool(item.Constant),
				Href:            item.Href,
				Query:           query,
			},
			Component: item.Component,
			Redirect:  "",
		}
		if item.Component == "view.iframe-page" {
			route.Props = map[string]any{
				"url": item.Href,
			}
			route.Meta.Href = ""
		}
		if item.IconType == "2" {
			route.Meta.LocalIcon = item.Icon
			route.Meta.Icon = ""
		}
		records = append(records, &route)
		uuidMap[item.Uuid] = &route
	}
	return records, uuidMap
}

func buildRouteTree(menus []*manage_menu.ManageMenu, uuidMap map[string]*types.Route, parentUuid string) []types.Route {
	var result []types.Route
	for _, menu := range menus {
		if menu.ParentUuid == parentUuid {
			if route, exists := uuidMap[menu.Uuid]; exists {
				subRoute := *route
				subRoute.Children = buildRouteTree(menus, uuidMap, menu.Uuid)
				result = append(result, subRoute)
			}
		}
	}
	return result
}
