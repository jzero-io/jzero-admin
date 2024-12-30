package route

import (
	"context"
	"net/http"
	"strings"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/route"
)

type GetConstantRoutes struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetConstantRoutes(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetConstantRoutes {
	return &GetConstantRoutes{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *GetConstantRoutes) GetConstantRoutes(req *types.GetConstantRoutesRequest) (resp []types.GetConstantRoutesResponseItem, err error) {
	menus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		Equal("constant", true).
		Build()...)
	if err != nil {
		return
	}
	resp = make([]types.GetConstantRoutesResponseItem, 0)
	for _, v := range menus {
		resp = append(resp, types.GetConstantRoutesResponseItem{
			Name:      v.RouteName,
			Path:      v.RoutePath,
			Component: v.Component,
			Props:     strings.Contains(v.RoutePath, ":"),
			Meta: types.RouteMeta{
				Title:      v.MenuName,
				I18nKey:    v.I18nKey,
				HideInMenu: cast.ToBool(v.HideInMenu),
				Constant:   cast.ToBool(v.Constant),
			},
		})
	}
	return
}
