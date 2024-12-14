package middleware

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Register(server *rest.Server) {
	httpx.SetOkHandler(ResponseMiddleware)
	httpx.SetErrorHandler(ErrorMiddleware)
	httpx.SetValidator(NewValidator())

	// i18n middleware
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			rctx := request.Context()
			if lang := request.Header.Get("Accept-Language"); lang == "" {
				rctx = context.WithValue(rctx, "lang", "zh-CN")
			} else {
				rctx = context.WithValue(rctx, "lang", lang)
			}
			next(writer, request.WithContext(rctx))
		}
	})
}
