package route

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/route"
)

type IsRouteExist struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsRouteExist(ctx context.Context, svcCtx *svc.ServiceContext) *IsRouteExist {
	return &IsRouteExist{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsRouteExist) IsRouteExist(req *types.IsRouteExistRequest) (resp *types.IsRouteExistResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
