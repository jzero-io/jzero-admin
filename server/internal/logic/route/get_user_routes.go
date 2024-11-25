package route

import (
	"context"
	"strings"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/auth"
	"server/internal/model/system_menu"
	"server/internal/svc"
	types "server/internal/types/route"
)

type GetUserRoutes struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRoutes(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRoutes {
	return &GetUserRoutes{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
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

	roleMenus, err := l.svcCtx.Model.SystemRoleMenu.FindByCondition(l.ctx, nil, condition.NewChain().
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

	menus, err := l.svcCtx.Model.SystemMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		In("id", uniqMenuIds).
		Build()...)
	if err != nil {
		return nil, err
	}
	list := buildRouteTree(convert(menus), 0)

	resp.Routes = list
	resp.Home = "home"
	return
}

func convert(list []*system_menu.SystemMenu) []*types.Route {
	var records []*types.Route
	for _, item := range list {
		var route types.Route
		route = types.Route{
			Id:       int64(item.Id),
			ParentId: item.ParentId,
			Name:     item.RouteName,
			Path:     item.RoutePath,
			Meta: types.RouteMeta{
				Title:      item.RouteName,
				I18nKey:    item.I18nKey,
				Icon:       item.Icon,
				Order:      int(item.Order),
				HideInMenu: cast.ToBool(item.HideInMenu),
				ActiveMenu: item.ActiveMenu.String,
				MultiTab:   cast.ToBool(item.MultiTab),
				KeepAlive:  cast.ToBool(item.KeepAlive),
				Constant:   cast.ToBool(item.Constant),
			},
			Component: item.Component,
			Props:     strings.Contains(item.RoutePath, ":"),
			Redirect:  "",
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
