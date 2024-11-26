package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/auth"
	"server/internal/svc"
	types "server/internal/types/auth"
	"server/pkg/jwt"
)

type PwdLogin struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPwdLogin(ctx context.Context, svcCtx *svc.ServiceContext) *PwdLogin {
	return &PwdLogin{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PwdLogin) PwdLogin(req *types.PwdLoginRequest) (resp *types.LoginResponse, err error) {
	user, err := l.svcCtx.Model.ManageUser.FindOneByUsername(l.ctx, nil, req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	if req.Password != user.Password {
		return nil, errors.New("用户名或密码错误")
	}
	userRoles, err := l.svcCtx.Model.ManageUserRole.FindByCondition(l.ctx, nil, condition.NewChain().
		Equal("user_id", user.Id).
		Build()...)
	if err != nil {
		return nil, err
	}
	var roleIds []int64
	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleId)
	}

	j := jwt.NewJwt(l.svcCtx.Config.Jwt.AccessSecret)
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
