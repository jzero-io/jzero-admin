package route

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/route"
)

type GetConstantRoutes struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConstantRoutes(ctx context.Context, svcCtx *svc.ServiceContext) *GetConstantRoutes {
	return &GetConstantRoutes{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConstantRoutes) GetConstantRoutes(req *types.GetConstantRoutesRequest) (resp *types.GetConstantRoutesResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
