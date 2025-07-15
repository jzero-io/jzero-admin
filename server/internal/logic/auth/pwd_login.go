package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/auth"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_user_role"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/auth"
)

func CreateToken(secret string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

type PwdLogin struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewPwdLogin(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *PwdLogin {
	return &PwdLogin{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *PwdLogin) PwdLogin(req *types.PwdLoginRequest) (resp *types.LoginResponse, err error) {
	config, err := l.svcCtx.Config.GetConfig()
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.Model.ManageUser.FindOneByUsername(l.ctx, nil, req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	if req.Password != user.Password {
		return nil, errors.New("用户名或密码错误")
	}
	userRoles, err := l.svcCtx.Model.ManageUserRole.FindByCondition(l.ctx, nil, condition.NewChain().
		Equal(manage_user_role.UserId, user.Id).
		Build()...)
	if err != nil {
		return nil, err
	}
	var roleIds []int64
	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleId)
	}

	marshal, err := json.Marshal(auth.Auth{
		Id:       int(user.Id),
		Username: user.Username,
		RoleIds:  roleIds,
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
	expirationTime := time.Now().Add(time.Duration(config.Jwt.AccessExpire) * time.Second).Unix()
	claims["exp"] = expirationTime

	token, err := CreateToken(l.svcCtx.MustGetConfig().Jwt.AccessSecret, claims)
	if err != nil {
		return nil, err
	}

	claims["exp"] = time.Now().Add(time.Duration(config.Jwt.RefreshExpire) * time.Second).Unix()
	refreshToken, err := CreateToken(l.svcCtx.MustGetConfig().Jwt.AccessSecret, claims)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
