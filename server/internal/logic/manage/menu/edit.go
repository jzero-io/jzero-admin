package menu

import (
	"context"
	"net/http"

	null "github.com/guregu/null/v5"
	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

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
	one, err := l.svcCtx.Model.ManageMenu.FindOne(l.ctx, nil, req.Id)
	if err != nil {
		return nil, err
	}
	var oldPermissions []types.Permission
	oldPermissionStr := one.Permissions.String

	one.Status = req.Status
	one.ParentId = int64(req.ParentId)
	one.MenuType = req.MenuType
	one.MenuName = req.MenuName
	one.HideInMenu = cast.ToInt64(req.HideInMenu)
	one.ActiveMenu = null.StringFrom(req.ActiveMenu).NullString
	one.Order = int64(req.Order)
	one.RouteName = req.RouteName
	one.RoutePath = req.RoutePath
	one.Component = req.Component
	one.Icon = req.Icon
	one.IconType = req.IconType
	one.I18nKey = req.I18nKey
	one.KeepAlive = cast.ToInt64(req.KeepAlive)
	one.Href = null.StringFrom(req.Href).NullString
	one.MultiTab = null.IntFrom(cast.ToInt64(req.MutiTab)).NullInt64
	one.FixedIndexInTab = null.IntFromPtr(req.FixedIndexInTab).NullInt64
	one.Query = null.StringFrom(marshal(req.Query)).NullString
	one.ButtonCode = null.StringFrom(req.ButtonCode).NullString
	one.Permissions = null.StringFrom(marshal(req.Permissions)).NullString
	one.Constant = cast.ToInt64(req.Constant)

	err = l.svcCtx.Model.ManageMenu.Update(l.ctx, nil, one)

	if req.MenuType == "2" || req.MenuType == "3" {
		// 更新了权限标识
		if marshal(req.Permissions) != oldPermissionStr {
			roleMenus, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.NewChain().
				Equal("menu_id", req.Id).
				Build()...)
			if err != nil {
				return nil, err
			}
			for _, rm := range roleMenus {
				// remove old casbin_rule for menu
				Unmarshal(oldPermissionStr, &oldPermissions)
				if len(oldPermissions) > 0 {
					for _, o := range oldPermissions {
						_, _ = l.svcCtx.CasbinEnforcer.RemovePolicy(cast.ToString(rm.RoleId), o.Code)
					}
				}

				// add casbin_rule
				var newPolicies [][]string
				permissions := req.Permissions

				for _, p := range permissions {
					newPolicies = append(newPolicies, []string{cast.ToString(rm.RoleId), p.Code})
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
