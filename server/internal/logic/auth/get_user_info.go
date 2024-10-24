package auth

import (
	"context"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/auth"
	"server/internal/svc"
	types "server/internal/types/auth"
)

type GetUserInfo struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfo(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfo {
	return &GetUserInfo{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfo) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	info, err := auth.Info(l.ctx)
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.Model.SystemUser.FindOne(l.ctx, uint64(info.Id))
	if err != nil {
		return nil, err
	}

	return &types.GetUserInfoResponse{
		UserId:   cast.ToString(user.Id),
		Username: user.Username,
		Roles:    []string{"R_super"},
		Buttons:  []string{"B_manage"},
	}, nil
}
