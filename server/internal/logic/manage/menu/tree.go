package menu

import (
	"context"
	"net/http"
	"sort"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/menu"
)

type Tree struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewTree(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Tree {
	return &Tree{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *Tree) Tree(req *types.TreeRequest) (resp []types.TreeResponse, err error) {
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
			sort.Slice(subMenu.Children, func(i, j int) bool {
				return subMenu.Children[i].Order < subMenu.Children[j].Order
			})
			result = append(result, subMenu)
		}
	}
	return result
}
