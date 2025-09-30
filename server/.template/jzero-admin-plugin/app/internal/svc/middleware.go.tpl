package svc

import (
    "net/http"

    "github.com/jzero-io/jzero-admin/core-engine/middleware"
)

type Middleware struct {
    middleware.Middleware
}

func NewMiddleware(svcCtx *ServiceContext, route2Code func(r *http.Request) string) Middleware {
	return Middleware{
    	Middleware: middleware.NewMiddleware(svcCtx.ServiceContext, route2Code),
    }
}