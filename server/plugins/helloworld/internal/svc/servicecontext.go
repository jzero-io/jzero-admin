package svc

import (
	"net/http"

	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-admin/core-engine/svc"
	"github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/stores/modelx"

	"helloworld/internal/config"
	"helloworld/internal/model"
)

type ServiceContext struct {
	*svc.ServiceContext
	ConfigCenter configcenter.ConfigCenter[config.Config]
	Model        model.Model
	Middleware
}

type ServiceContextOpts struct {
	CoreServiceContext *svc.ServiceContext
}

func (opts ServiceContextOpts) DefaultOptions() ServiceContextOpts {
	return ServiceContextOpts{}
}

func WithCoreServiceContext(coreSvcCtx *svc.ServiceContext) opts.Opt[ServiceContextOpts] {
	return func(opts *ServiceContextOpts) {
		opts.CoreServiceContext = coreSvcCtx
	}
}

func NewServiceContext(cc configcenter.ConfigCenter[config.Config], route2code func(r *http.Request) string, op ...opts.Opt[ServiceContextOpts]) *ServiceContext {
	o := opts.DefaultApply(op...)
	svcCtx := &ServiceContext{
		ConfigCenter: cc,
	}
	svcCtx.SetConfigListener()

	if o.CoreServiceContext != nil {
		svcCtx.ServiceContext = o.CoreServiceContext
	} else {
		svcCtx.ServiceContext = svc.NewServiceContext(svcCtx.MustGetConfig().Config, route2code)
	}

	svcCtx.Model = model.NewModel(svcCtx.SqlxConn, modelx.WithCachedConn(modelx.NewConnWithCache(svcCtx.SqlxConn, svcCtx.Cache)))
	svcCtx.Middleware = NewMiddleware(svcCtx)
	return svcCtx
}
