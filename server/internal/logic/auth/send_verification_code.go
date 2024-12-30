package auth

import (
	"context"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	gomail "gopkg.in/gomail.v2"

	"github.com/jzero-io/jzero-admin/server/internal/constant"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/auth"
)

var SendVerificationError = errors.New("发送失败, 请联系管理员")

type SendVerificationCode struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewSendVerificationCode(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *SendVerificationCode {
	return &SendVerificationCode{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *SendVerificationCode) SendVerificationCode(req *types.SendVerificationCodeRequest) (resp *types.SendVerificationCodeResponse, err error) {
	if req.VerificationType == "email" {
		email, err := l.svcCtx.Model.ManageEmail.FindOneByCondition(l.ctx, nil)
		if err != nil {
			return nil, SendVerificationError
		}

		verificationUuid := uuid.New().String()
		verificationCode := genValidateCode(6)

		m := gomail.NewMessage()
		// 设置发件人
		m.SetHeader("From", email.From)
		// 设置收件人
		m.SetHeader("To", req.Email)
		// 设置邮件主题
		m.SetHeader("Subject", "验证码")
		// 设置邮件正文
		m.SetBody("text/plain", fmt.Sprintf("JzeroAdmin 邮箱验证码: %s", verificationCode))
		// 配置 SMTP 服务器信息
		d := gomail.NewDialer(email.Host, cast.ToInt(email.Port), email.Username, email.Password)
		d.SSL = cast.ToBool(email.EnableSsl)
		if !cast.ToBool(email.IsVerify) {
			tlsConfig := &tls.Config{
				InsecureSkipVerify: true,
			}
			d.TLSConfig = tlsConfig
		}
		// 发送邮件
		if err = d.DialAndSend(m); err != nil {
			return nil, SendVerificationError
		}

		if err = l.svcCtx.Cache.SetWithExpireCtx(context.Background(), fmt.Sprintf("%s:%s", constant.CacheVerificationCodePrefix, verificationUuid), verificationCode, time.Minute*5); err != nil {
			return nil, SendVerificationError
		}

		var cacheVal string
		if err = l.svcCtx.Cache.Get(fmt.Sprintf("%s:%s", constant.CacheVerificationCodePrefix, verificationUuid), &cacheVal); err == nil {
			logx.Infof("get cache %s:%s", verificationUuid, cacheVal)
		}
		return &types.SendVerificationCodeResponse{
			VerificationUuid: verificationUuid,
		}, err
	}
	return nil, errors.New("暂不支持手机号验证码")
}

func genValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder

	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
