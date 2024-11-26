package menu

import (
	"context"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/manage/menu"
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

func (l *GetAllPages) GetAllPages(req *types.GetAllPagesRequest) (resp []string, err error) {
	pages, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		Equal("menu_type", "2").
		Equal("status", "1").
		NotEqual("hide_in_menu", true).
		Build()...)
	if err != nil {
		return
	}
	for _, page := range pages {
		resp = append(resp, page.RoutePath)
	}
	return
}
