package middleware

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"helloworld/internal/global"
)

func Register(server *rest.Server) {
	httpx.SetOkHandler(global.ServiceContext.Ok)
	httpx.SetErrorHandlerCtx(global.ServiceContext.Error)
	httpx.SetValidator(global.ServiceContext.Validate)
	server.Use(global.ServiceContext.I18n)
}
