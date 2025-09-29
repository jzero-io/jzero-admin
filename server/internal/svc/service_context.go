package svc

import (
	"net/http"

	"github.com/jzero-io/jzero-admin/core-engine/svc"
	"github.com/jzero-io/jzero/core/stores/modelx"
	configurator "github.com/zeromicro/go-zero/core/configcenter"

	"github.com/jzero-io/jzero-admin/server/internal/config"
	"github.com/jzero-io/jzero-admin/server/internal/model"
)

type ServiceContext struct {
	*svc.ServiceContext
	Config configurator.Configurator[config.Config]
	Model  model.Model
	Middleware
}

func NewServiceContext(cc configurator.Configurator[config.Config], route2code func(r *http.Request) string) *ServiceContext {
	svcCtx := &ServiceContext{
		Config: cc,
	}
	svcCtx.SetConfigListener()

	svcCtx.ServiceContext = svc.NewServiceContext(svcCtx.MustGetConfig().Config, route2code)
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn, modelx.WithCachedConn(modelx.NewConnWithCache(svcCtx.SqlxConn, svcCtx.Cache)))
	svcCtx.Middleware = NewMiddleware(svcCtx, route2code)
	return svcCtx
}
