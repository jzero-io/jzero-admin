package route

import (
	"context"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_menu"
	"github.com/jzero-io/jzero/core/stores/condition"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/route"
)

type IsRouteExist struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewIsRouteExist(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *IsRouteExist {
	return &IsRouteExist{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *IsRouteExist) IsRouteExist(req *types.IsRouteExistRequest) (resp bool, err error) {
	manageMenu, err := l.svcCtx.Model.ManageMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
		Equal(manage_menu.ManageMenuField.RouteName, req.RouteName).
		Build()...)

	return manageMenu != nil, err
}
