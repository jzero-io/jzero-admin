package svc

type Middleware struct {
}

func NewMiddleware(svcCtx *ServiceContext) Middleware {
	return Middleware{}
}