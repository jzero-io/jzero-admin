package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jzero-io/jzero/core/status"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"

	"github.com/jzero-io/jzero-admin/server/internal/errcodes/auth"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/auth"
)

var (
	RefreshTokenExpiredErr = errors.New("refresh token expired")
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
	config, err := l.svcCtx.ConfigCenter.GetConfig()
	if err != nil {
		return nil, err
	}

	// 解析 refreshToken
	parser := token.NewTokenParser()
	tok, err := parser.ParseToken(&http.Request{
		Header: http.Header{
			"Authorization": []string{req.RefreshToken},
		},
	}, l.svcCtx.MustGetConfig().Jwt.AccessSecret, "")
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, status.Wrap(auth.RefreshTokenExpiredCode, RefreshTokenExpiredErr)
		}
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// 设置新的过期时间
	claims["exp"] = time.Now().Add(time.Duration(config.Jwt.AccessExpire) * time.Second).Unix()
	newAccessToken, err := CreateToken(l.svcCtx.MustGetConfig().Jwt.AccessSecret, claims)
	if err != nil {
		return nil, err
	}

	claims["exp"] = time.Now().Add(time.Duration(config.Jwt.RefreshExpire) * time.Second).Unix()
	newRefreshToken, err := CreateToken(l.svcCtx.MustGetConfig().Jwt.AccessSecret, claims)
	if err != nil {
		return nil, err
	}

	return &types.RefreshTokenResponse{
		Token:        newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
