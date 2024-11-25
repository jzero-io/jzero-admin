package menu

import (
	"context"
	"encoding/json"
	"sort"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/model/system_menu"
	"server/internal/svc"
	types "server/internal/types/system/menu"
)

type List struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewList(ctx context.Context, svcCtx *svc.ServiceContext) *List {
	return &List{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *List) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	resp = &types.ListResponse{}

	list, err := l.svcCtx.Model.SystemMenu.FindByCondition(l.ctx, nil)
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

func unmarshal(marshalStr string, v any) {
	_ = json.Unmarshal([]byte(marshalStr), v)
}

func convert(list []*system_menu.SystemMenu) []*types.SystemMenu {
	var records []*types.SystemMenu
	for _, item := range list {
		var menu types.SystemMenu
		unmarshal(item.Buttons.String, &menu.Buttons)
		menu = types.SystemMenu{
			Id:         item.Id,
			ParentId:   uint64(item.ParentId),
			MenuType:   item.MenuType,
			MenuName:   item.MenuName,
			RouteName:  item.RouteName,
			RoutePath:  item.RoutePath,
			Component:  item.Component,
			Icon:       item.Icon,
			IconType:   item.IconType,
			Order:      uint64(item.Order),
			I18nKey:    item.I18nKey,
			Status:     item.Status,
			Constant:   cast.ToBool(item.Constant),
			HideInMenu: cast.ToBool(item.HideInMenu),
			MultiTab:   cast.ToBool(item.MultiTab),
			KeepAlive:  cast.ToBool(item.KeepAlive),
			ActiveMenu: item.ActiveMenu.String,
			Children:   nil,
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
