package auth

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/auth"
)

type Error struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewError(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Error {
	return &Error{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *Error) Error(req *types.ErrorRequest) (resp *types.ErrorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
