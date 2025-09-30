package middleware

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"helloworld/internal/global"
)

func Register(server *rest.Server) {
	httpx.SetOkHandler(global.ServiceContext.Middleware.Ok)
	httpx.SetErrorHandlerCtx(global.ServiceContext.Middleware.Error)
	httpx.SetValidator(global.ServiceContext.Middleware.Validate)
	server.Use(global.ServiceContext.Middleware.I18n)
}
