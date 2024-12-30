package svc

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/server/internal/middleware"
)

type Middleware struct {
	Authx rest.Middleware
}

func NewMiddleware(svcCtx *ServiceContext, route2Code func(r *http.Request) string) Middleware {
	return Middleware{
		Authx: middleware.NewAuthxMiddleware(svcCtx.CasbinEnforcer, route2Code).Handle,
	}
}
