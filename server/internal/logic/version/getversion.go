package version

import (
	"context"

	"server/internal/svc"
	types "server/internal/types/version"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVersion struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVersion(ctx context.Context, svcCtx *svc.ServiceContext) *GetVersion {
	return &GetVersion{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVersion) GetVersion(req *types.GetVersionRequest) (resp *types.GetVersionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
