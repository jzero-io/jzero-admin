package middleware

import (
	casbin "github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Middleware struct {
	CasbinEnforcer *casbin.Enforcer

	Authx rest.Middleware
}

func NewMiddleware(casbinEnforcer *casbin.Enforcer, route2Code func(r *http.Request) string) Middleware {
	return Middleware{
		CasbinEnforcer: casbinEnforcer,
		Authx:          NewAuthxMiddleware(casbinEnforcer, route2Code).Handle,
	}
}

func Register(server *rest.Server) {
	httpx.SetOkHandler(ResponseMiddleware)
	httpx.SetErrorHandler(ErrorMiddleware)
	httpx.SetValidator(NewValidator())
}
