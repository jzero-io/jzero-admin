package role

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/role"
)

type GetMenus struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetMenus(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetMenus {
	return &GetMenus{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *GetMenus) GetMenus(req *types.GetMenusRequest) (resp []uint64, err error) {
	menus, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.Condition{
		Field:    "role_id",
		Operator: condition.Equal,
		Value:    req.RoleId,
	})
	if err != nil {
		return
	}

	var menuIds []uint64
	for _, menu := range menus {
		menuIds = append(menuIds, uint64(menu.MenuId))
	}
	resp = menuIds

	return
}
