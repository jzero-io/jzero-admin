package auth

import (
	"context"

	"github.com/pkg/errors"
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
	user, err := l.svcCtx.Model.SystemUser.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	if req.Password != user.Password {
		return nil, errors.New("用户名或密码错误")
	}

	return &types.LoginResponse{
		Token:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
	}, nil
}
