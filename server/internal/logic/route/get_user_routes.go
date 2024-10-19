package route

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/route"
)

type GetUserRoutes struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRoutes(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRoutes {
	return &GetUserRoutes{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRoutes) GetUserRoutes(req *types.GetUserRoutesRequest) (resp *types.GetUserRoutesResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
