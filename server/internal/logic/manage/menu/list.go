package menu

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/jzero-io/jzero-contrib/nullx"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_menu"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/menu"
)

type List struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewList(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *List {
	return &List{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *List) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	resp = &types.ListResponse{}

	list, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		Equal("constant", false).
		Build()...)
	if err != nil {
		return nil, err
	}

	tree := buildMenuTree(convert(list), 0)

	// sort by order asc
	sort.Slice(tree, func(i, j int) bool {
		return tree[i].Order < tree[j].Order
	})

	// page by current and size
	resp.PageResponse = types.PageResponse{
		Current: req.Current,
		Size:    req.Size,
		Total:   int64(len(tree)),
	}
	resp.Records = paginate(tree, req.Current, req.Size)

	return
}

func Unmarshal(marshalStr string, v any) {
	_ = json.Unmarshal([]byte(marshalStr), v)
}

func convert(list []*manage_menu.ManageMenu) []*types.SystemMenu {
	var records []*types.SystemMenu
	for _, item := range list {
		var menu types.SystemMenu
		var permissions []types.Permission
		var query []types.Query
		Unmarshal(item.Permissions.String, &permissions)
		Unmarshal(item.Query.String, &query)
		menu = types.SystemMenu{
			Id:              item.Id,
			ActiveMenu:      item.ActiveMenu.String,
			MenuType:        item.MenuType,
			MenuName:        item.MenuName,
			RouteName:       item.RouteName,
			RoutePath:       item.RoutePath,
			Component:       item.Component,
			Icon:            item.Icon,
			IconType:        item.IconType,
			ParentId:        uint64(item.ParentId),
			Status:          item.Status,
			KeepAlive:       cast.ToBool(item.KeepAlive),
			Constant:        cast.ToBool(item.Constant),
			Order:           uint64(item.Order),
			HideInMenu:      cast.ToBool(item.HideInMenu),
			Href:            item.Href.String,
			MultiTab:        cast.ToBool(item.MultiTab),
			FixedIndexInTab: nullx.NewInt(item.FixedIndexInTab).Ptr(),
			Query:           query,
			ButtonCode:      item.ButtonCode.String,
			Permissions:     permissions,
			I18nKey:         item.I18nKey,
			Children:        nil,
		}
		records = append(records, &menu)
	}
	return records
}

func buildMenuTree(menus []*types.SystemMenu, parentId uint64) []types.SystemMenu {
	var result []types.SystemMenu
	for _, menu := range menus {
		if menu.ParentId == parentId {
			subMenu := *menu
			subMenu.Children = buildMenuTree(menus, menu.Id)
			sort.Slice(subMenu.Children, func(i, j int) bool {
				return subMenu.Children[i].Order < subMenu.Children[j].Order
			})
			result = append(result, subMenu)
		}
	}
	return result
}

func paginate(list []types.SystemMenu, current, size int) []types.SystemMenu {
	start := (current - 1) * size
	if start >= len(list) {
		return []types.SystemMenu{}
	}

	end := start + size
	if end > len(list) {
		end = len(list)
	}

	return list[start:end]
}
