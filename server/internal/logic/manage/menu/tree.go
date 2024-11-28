package menu

import (
	"context"
	"sort"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/manage/menu"
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

func (l *Tree) Tree(req *types.TreeRequest) (resp []types.TreeResponse, err error) {
	resp = []types.TreeResponse{}

	list, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		NotEqual("constant", true).
		Build()...)
	if err != nil {
		return nil, err
	}

	tree := buildSimpleMenuTree(convert(list), 0)

	// sort by order asc
	sort.Slice(tree, func(i, j int) bool {
		return tree[i].Order < tree[j].Order
	})

	resp = tree

	return
}

func buildSimpleMenuTree(menus []*types.SystemMenu, parentId uint64) []types.TreeResponse {
	var result []types.TreeResponse
	for _, menu := range menus {
		if menu.ParentId == parentId {
			subMenu := types.TreeResponse{
				Id:    menu.Id,
				Label: menu.MenuName,
				PId:   menu.ParentId,
				Order: menu.Order,
			}
			subMenu.Children = buildSimpleMenuTree(menus, menu.Id)
			result = append(result, subMenu)
		}
	}
	return result
}
