package route

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

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
	resp = &types.GetUserRoutesResponse{}

	list := []types.Route{
		{
			Name:      "home",
			Path:      "/home",
			Component: "layout.base$view.home",
			Meta: types.RouteMeta{
				Title:   "home",
				I18nKey: "route.home",
				Icon:    "mdi:monitor-dashboard",
				Order:   1,
			},
		},
		{
			Name:      "about",
			Path:      "/about",
			Component: "layout.base$view.about",
			Meta: types.RouteMeta{
				Title:   "about",
				I18nKey: "route.about",
				Icon:    "fluent:book-information-24-regular",
				Order:   10,
			},
		},
		{
			Name:      "manage",
			Path:      "/manage",
			Component: "layout.base",
			Meta: types.RouteMeta{
				I18nKey: "route.manage",
				Icon:    "carbon:cloud-service-management",
				Order:   9,
				Title:   "manage",
			},
			Children: []types.Route{
				{
					Name:      "manage_menu",
					Path:      "/manage/menu",
					Component: "view.manage_menu",
					Meta: types.RouteMeta{
						I18nKey:   "route.manage_menu",
						Icon:      "material-symbols:route",
						Order:     3,
						KeepAlive: true,
						Title:     "manage_menu",
					},
				},
				{
					Name:      "manage_user",
					Path:      "/manage/user",
					Component: "view.manage_user",
					Meta: types.RouteMeta{
						I18nKey: "route.manage_user",
						Icon:    "ic:round-manage-accounts",
						Order:   1,
						Title:   "manage_user",
					},
				},
				{
					Name:      "manage_user-detail",
					Path:      "/manage/user-detail/:id",
					Component: "view.manage_user-detail",
					Meta: types.RouteMeta{
						I18nKey:    "route.manage_user-detail",
						Icon:       "ic:round-manage-accounts",
						Title:      "manage_user-detail",
						HideInMenu: true,
						ActiveMenu: "manage_user",
					},
				},
				{
					Path:      "/manage/role",
					Component: "view.manage_role",
					Name:      "manage_role",
					Meta: types.RouteMeta{
						I18nKey: "route.manage_role",
						Icon:    "carbon:user-role",
						Order:   2,
						Title:   "manage_role",
					},
				},
			},
		},
	}
	resp.Routes = list
	resp.Home = "home"
	return
}
