package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/system_manage/role"
)

type GetAll struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAll(ctx context.Context, svcCtx *svc.ServiceContext) *GetAll {
	return &GetAll{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAll) GetAll(req *types.GetAllRequest) (resp *types.GetAllResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
