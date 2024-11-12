package role

import (
	"context"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/system/role"
)

type GetMenus struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenus(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenus {
	return &GetMenus{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenus) GetMenus(req *types.GetMenusRequest) (resp []uint64, err error) {
	menus, err := l.svcCtx.Model.SystemRoleMenu.FindByCondition(l.ctx, condition.Condition{
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
