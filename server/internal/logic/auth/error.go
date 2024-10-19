package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/auth"
)

type Error struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewError(ctx context.Context, svcCtx *svc.ServiceContext) *Error {
	return &Error{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Error) Error(req *types.ErrorRequest) (resp *types.ErrorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
