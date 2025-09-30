package svc

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/core-engine/middleware"
)

type Middleware struct {
	Authx    rest.Middleware
	Ok       func(ctx context.Context, data any) any
	Error    func(ctx context.Context, err error) (int, any)
	I18n     rest.Middleware
	Validate *middleware.ValidatorMiddleware
}

func NewMiddleware(svcCtx *ServiceContext, route2code func(r *http.Request) string) Middleware {
	return Middleware{
		Error:    middleware.NewErrorMiddleware().Handle,
		Ok:       middleware.NewOkMiddleware().Handle,
		Authx:    middleware.NewAuthxMiddleware(svcCtx.CasbinEnforcer, route2code).Handle,
		I18n:     middleware.NewI18nMiddleware().Handle,
		Validate: middleware.NewValidatorMiddleware(),
	}
}
