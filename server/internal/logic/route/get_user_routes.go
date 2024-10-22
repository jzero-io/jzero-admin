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
			Children: []types.Route{
				{
					Meta: types.RouteMeta{
						I18nKey:   "route.manage_menu",
						Icon:      "material-symbols:route",
						Order:     3,
						KeepAlive: true,
						Title:     "manage_menu",
					},
					Name:      "manage_menu",
					Path:      "/manage/menu",
					Component: "view.manage_menu",
				},
				{
					Name: "manage_role",
					Meta: types.RouteMeta{
						I18nKey: "route.manage_role",
						Icon:    "carbon:user-role",
						Order:   2,
						Title:   "manage_role",
					},
					Path:      "/manage/role",
					Component: "view.manage_role",
				},
			},
			Meta: types.RouteMeta{
				I18nKey: "route.manage",
				Icon:    "carbon:cloud-service-management",
				Order:   9,
				Title:   "manage",
			},
			Name:      "manage",
			Path:      "/manage",
			Component: "view.manage_menu",
		},
	}
	resp.Routes = list
	return
}
