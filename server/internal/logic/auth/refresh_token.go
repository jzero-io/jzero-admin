package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/auth"
	"github.com/jzero-io/jzero-admin/server/pkg/jwt"
)

type RefreshToken struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewRefreshToken(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *RefreshToken {
	return &RefreshToken{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *RefreshToken) RefreshToken(req *types.RefreshTokenRequest) (resp *types.RefreshTokenResponse, err error) {
	config, err := l.svcCtx.Config.GetConfig()
	if err != nil {
		return nil, err
	}

	// 解析 refreshToken
	j := jwt.NewJwt(config.Jwt.AccessSecret)
	claims, err := j.ParseToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("无效的 refreshToken")
	}

	// 验证 refreshToken 是否过期
	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return nil, errors.New("refreshToken 已过期")
	}

	// 设置新的过期时间
	claims["exp"] = time.Now().Add(time.Duration(config.Jwt.AccessExpire) * time.Second).Unix()
	newAccessToken, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}

	claims["exp"] = time.Now().Add(time.Duration(config.Jwt.RefreshExpire) * time.Second).Unix()
	newRefreshToken, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}

	return &types.RefreshTokenResponse{
		Token:        newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
