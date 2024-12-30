package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/guregu/null/v5"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/constant"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_user"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/auth"
)

var RegisterError = errors.New("注册失败")

type Register struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewRegister(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Register {
	return &Register{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
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

	_, err = l.svcCtx.Model.ManageUser.FindOneByUsername(l.ctx, nil, req.Username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	_, err = l.svcCtx.Model.ManageUser.Insert(l.ctx, nil, &manage_user.ManageUser{
		Username:   req.Username,
		Password:   req.Password,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Email:      null.StringFrom(req.Email).NullString,
		Gender:     "1",
		Status:     "1",
	})

	return
}
