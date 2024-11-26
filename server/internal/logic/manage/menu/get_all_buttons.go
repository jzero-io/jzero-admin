package menu

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/manage/menu"
)

type GetAllButtons struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllButtons(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllButtons {
	return &GetAllButtons{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllButtons) GetAllButtons(req *types.Empty) (resp []types.Button, err error) {
	// todo: add your logic here and delete this line

	return
}
