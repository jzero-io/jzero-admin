package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/guregu/null/v5"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/constant"
	"server/internal/model/system_user"
	"server/internal/svc"
	types "server/internal/types/auth"
)

var RegisterError = errors.New("注册失败")

type Register struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegister(ctx context.Context, svcCtx *svc.ServiceContext) *Register {
	return &Register{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Register) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// check verificationUuid
	var verificationUuidVal string
	if err = l.svcCtx.Cache.Get(fmt.Sprintf("%s:%s", constant.CacheVerificationCodePrefix, req.VerificationUuid), &verificationUuidVal); err != nil {
		return nil, RegisterError
	}
	if verificationUuidVal != req.VerificationCode {
		return nil, errors.New("验证码错误")
	}

	_, err = l.svcCtx.Model.SystemUser.FindOneByUsername(l.ctx, req.Username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	_, err = l.svcCtx.Model.SystemUser.Insert(l.ctx, &system_user.SystemUser{
		Username:   req.Username,
		Password:   req.Password,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Email:      null.StringFrom(req.Email).NullString,
	})

	return
}
