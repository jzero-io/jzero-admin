package menu

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/pkg/errors"
	"github.com/samber/lo"
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
	data, err := l.svcCtx.Model.ManageMenu.FindOneByUuid(l.ctx, nil, req.Uuid)
	if err != nil {
		return nil, err
	}
	var oldPermissions []types.Permission

	newData := lo.FromPtr(data)
	newData.Status = req.Status
	newData.ParentUuid = req.ParentUuid
	newData.MenuType = req.MenuType
	newData.MenuName = req.MenuName
	newData.HideInMenu = cast.ToInt64(req.HideInMenu)
	newData.ActiveMenu = req.ActiveMenu
	newData.Order = req.Order
	newData.RouteName = req.RouteName
	newData.RoutePath = req.RoutePath
	newData.Component = req.Component
	newData.Icon = req.Icon
	newData.IconType = req.IconType
	newData.I18nKey = req.I18nKey
	newData.KeepAlive = cast.ToInt64(req.KeepAlive)
	newData.Href = req.Href
	newData.MultiTab = cast.ToInt64(req.MutiTab)
	newData.FixedIndexInTab = cast.ToInt64(req.FixedIndexInTab)
	newData.Query = marshal(req.Query)
	newData.ButtonCode = req.ButtonCode
	newData.Permissions = marshal(req.Permissions)
	newData.Constant = cast.ToInt64(req.Constant)

	if err = l.svcCtx.Model.ManageMenu.Update(l.ctx, nil, lo.ToPtr(newData)); err != nil {
		return nil, err
	}

	if req.MenuType == "2" || req.MenuType == "3" {
		// 更新了权限标识
		if marshal(req.Permissions) != data.Permissions {
			roleMenus, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.NewChain().
				Equal(manage_role_menu.MenuUuid, req.Uuid).
				Build()...)
			if err != nil {
				return nil, err
			}
			for _, rm := range roleMenus {
				// remove old casbin_rule for menu
				Unmarshal(data.Permissions, &oldPermissions)
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
