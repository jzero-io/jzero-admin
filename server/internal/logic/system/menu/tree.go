package menu

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/system/menu"
)

type Tree struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTree(ctx context.Context, svcCtx *svc.ServiceContext) *Tree {
	return &Tree{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Tree) Tree(req *types.TreeRequest) (resp *types.TreeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
