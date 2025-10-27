package menu

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role_menu"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/menu"
)

type Edit struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewEdit(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Edit {
	return &Edit{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *Edit) Edit(req *types.EditRequest) (resp *types.EditResponse, err error) {
	one, err := l.svcCtx.Model.ManageMenu.FindOneByUuid(l.ctx, nil, req.Uuid)
	if err != nil {
		return nil, err
	}
	var oldPermissions []types.Permission
	oldPermissionStr := one.Permissions

	one.Status = req.Status
	one.ParentUuid = req.ParentUuid
	one.MenuType = req.MenuType
	one.MenuName = req.MenuName
	one.HideInMenu = cast.ToInt64(req.HideInMenu)
	one.ActiveMenu = req.ActiveMenu
	one.Order = req.Order
	one.RouteName = req.RouteName
	one.RoutePath = req.RoutePath
	one.Component = req.Component
	one.Icon = req.Icon
	one.IconType = req.IconType
	one.I18nKey = req.I18nKey
	one.KeepAlive = cast.ToInt64(req.KeepAlive)
	one.Href = req.Href
	one.MultiTab = cast.ToInt64(req.MutiTab)
	one.FixedIndexInTab = cast.ToInt64(req.FixedIndexInTab)
	one.Query = marshal(req.Query)
	one.ButtonCode = req.ButtonCode
	one.Permissions = marshal(req.Permissions)
	one.Constant = cast.ToInt64(req.Constant)

	err = l.svcCtx.Model.ManageMenu.Update(l.ctx, nil, one)

	if req.MenuType == "2" || req.MenuType == "3" {
		// 更新了权限标识
		if marshal(req.Permissions) != oldPermissionStr {
			roleMenus, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.NewChain().
				Equal(manage_role_menu.ManageRoleMenuField.MenuUuid, req.Uuid).
				Build()...)
			if err != nil {
				return nil, err
			}
			for _, rm := range roleMenus {
				// remove old casbin_rule for menu
				Unmarshal(oldPermissionStr, &oldPermissions)
				if len(oldPermissions) > 0 {
					for _, o := range oldPermissions {
						_, _ = l.svcCtx.CasbinEnforcer.RemovePolicy(rm.RoleUuid, o.Code)
					}
				}

				// add casbin_rule
				var newPolicies [][]string
				permissions := req.Permissions

				for _, p := range permissions {
					newPolicies = append(newPolicies, []string{rm.RoleUuid, p.Code})
				}

				if len(newPolicies) > 0 {
					var b bool
					b, _ = l.svcCtx.CasbinEnforcer.AddPolicies(newPolicies)
					if !b {
						return nil, errors.Wrapf(err, "failed to add policies")
					}
					err = l.svcCtx.CasbinEnforcer.LoadPolicy()
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}
	return
}
