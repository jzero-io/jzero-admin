package middleware

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/core-engine/svc"
)

type Middleware struct {
	Authx    rest.Middleware
	Ok       func(ctx context.Context, data any) any
	Error    func(ctx context.Context, err error) (int, any)
	I18n     rest.Middleware
	Validate *ValidatorMiddleware
}

func NewMiddleware(svcCtx *svc.ServiceContext, route2code func(r *http.Request) string) Middleware {
	return Middleware{
		Error:    NewErrorMiddleware().Handle,
		Ok:       NewOkMiddleware().Handle,
		Authx:    NewAuthxMiddleware(svcCtx.CasbinEnforcer, route2code).Handle,
		I18n:     NewI18nMiddleware().Handle,
		Validate: NewValidatorMiddleware(),
	}
}
