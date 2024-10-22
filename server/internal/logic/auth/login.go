package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/auth"
)

type Login struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogin(ctx context.Context, svcCtx *svc.ServiceContext) *Login {
	return &Login{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Login) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	return &types.LoginResponse{
		Token:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
	}, nil
}
