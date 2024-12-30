package menu

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/jzero-io/jzero-contrib/status"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/errcodes/manage"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/menu"
)

type Delete struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewDelete(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Delete {
	return &Delete{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *Delete) Delete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	if len(req.Ids) == 0 {
		return nil, nil
	}
	// whether it has submenu
	subMenus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.Condition{
		Field:    "parent_id",
		Operator: condition.In,
		Value:    req.Ids,
	})
	if err == nil && len(subMenus) > 0 {
		return nil, status.ErrorMessage(manage.ExistSubMenuCode, l.svcCtx.Trans.Trans(l.ctx, "manage.menu.existSubMenu"))
	}
	// remove permissions
	menus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.Condition{
		Field:    "id",
		Operator: condition.In,
		Value:    req.Ids,
	})
	if err == nil {
		for _, menu := range menus {
			var permissions []types.Permission
			Unmarshal(menu.Permissions.String, &permissions)
			if len(permissions) > 0 {
				roles, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.Condition{
					Field:    "menu_id",
					Operator: condition.Equal,
					Value:    menu.Id,
				})
				if err != nil {
					return nil, err
				}
				for _, role := range roles {
					for _, permission := range permissions {
						_, _ = l.svcCtx.CasbinEnforcer.RemovePolicy(cast.ToString(role.RoleId), permission.Code)
					}
				}
			}
		}
	}
	err = l.svcCtx.Model.ManageMenu.DeleteByCondition(l.ctx, nil, condition.Condition{
		Field:    "id",
		Operator: condition.In,
		Value:    req.Ids,
	})
	return
}
