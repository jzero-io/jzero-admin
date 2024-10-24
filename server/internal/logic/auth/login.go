package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/auth"
	"server/internal/svc"
	types "server/internal/types/auth"
	"server/pkg/jwt"
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

	j := jwt.NewJwt(l.svcCtx.Config.Jwt.AccessSecret)
	marshal, err := json.Marshal(auth.Auth{
		Id:       int(user.Id),
		Username: user.Username,
	})
	if err != nil {
		return nil, err
	}

	var claims map[string]any
	err = json.Unmarshal(marshal, &claims)
	if err != nil {
		return nil, err
	}

	// token 过期时间
	expirationTime := time.Now().Add(time.Duration(l.svcCtx.Config.Jwt.AccessExpire) * time.Second).Unix()
	claims["exp"] = expirationTime

	token, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}

	claims["exp"] = time.Now().Add(time.Duration(l.svcCtx.Config.Jwt.RefreshExpire) * time.Second).Unix()
	refreshToken, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
