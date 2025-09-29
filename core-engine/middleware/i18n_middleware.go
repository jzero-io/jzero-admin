package middleware

import (
	"context"
	"net/http"
)

type I18nMiddleware struct {
}

func NewI18nMiddleware() *I18nMiddleware {
	return &I18nMiddleware{}
}

func (m *I18nMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		rctx := request.Context()
		if lang := request.Header.Get("Accept-Language"); lang == "" {
			rctx = context.WithValue(rctx, "lang", "zh-CN")
		} else {
			rctx = context.WithValue(rctx, "lang", lang)
		}
		next(writer, request.WithContext(rctx))
	}
}
