package version

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/version"
)

type Get struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet(ctx context.Context, svcCtx *svc.ServiceContext) *Get {
	return &Get{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get) Get(req *types.GetRequest) (resp *types.GetResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
