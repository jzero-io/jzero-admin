package user

import (
	"context"
	"time"

	"github.com/guregu/null/v5"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/system/user"
)

type Edit struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEdit(ctx context.Context, svcCtx *svc.ServiceContext) *Edit {
	return &Edit{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Edit) Edit(req *types.EditRequest) (resp *types.EditResponse, err error) {
	user, err := l.svcCtx.Model.SystemUser.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	user.Username = req.Username
	user.Nickname = req.NickName
	user.Email = null.StringFrom(req.UserEmail).NullString
	user.Phone = null.StringFrom(req.UserPhone).NullString
	user.Gender = req.UserGender
	user.Status = req.Status
	user.UpdateTime = time.Now()

	err = l.svcCtx.Model.SystemUser.Update(l.ctx, user)
	return
}
