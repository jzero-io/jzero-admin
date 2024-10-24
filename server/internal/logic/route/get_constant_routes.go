package route

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/route"
)

type GetConstantRoutes struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConstantRoutes(ctx context.Context, svcCtx *svc.ServiceContext) *GetConstantRoutes {
	return &GetConstantRoutes{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConstantRoutes) GetConstantRoutes(req *types.GetConstantRoutesRequest) (resp []types.GetConstantRoutesResponseItem, err error) {
	resp = []types.GetConstantRoutesResponseItem{
		{
			Name:      "login",
			Path:      "/login/:module(pwd-login|code-login|register|reset-pwd|bind-wechat)?",
			Component: "layout.blank$view.login",
			Props:     true,
			Meta: types.RouteMeta{
				Title:      "login",
				I18nKey:    "route.login",
				HideInMenu: true,
				Constant:   true,
			},
		},
		{
			Name:      "403",
			Path:      "/403",
			Component: "layout.blank$view.403",
			Meta: types.RouteMeta{
				Title:      "403",
				I18nKey:    "route.403",
				HideInMenu: true,
				Constant:   true,
			},
		},
		{
			Name:      "404",
			Path:      "/404",
			Component: "layout.blank$view.404",
			Meta: types.RouteMeta{
				Title:      "404",
				I18nKey:    "route.404",
				HideInMenu: true,
				Constant:   true,
			},
		},
		{
			Name:      "500",
			Path:      "/500",
			Component: "layout.blank$view.500",
			Meta: types.RouteMeta{
				Title:      "500",
				I18nKey:    "route.500",
				HideInMenu: true,
				Constant:   true,
			},
		},
	}
	return
}
