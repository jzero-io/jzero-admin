package menu

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/system_manage/menu"
)

type GetAllPages struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPages(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPages {
	return &GetAllPages{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPages) GetAllPages(req *types.GetAllPagesRequest) (resp *types.GetAllPagesResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
