package role

import (
	"context"
	"time"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/logic/manage/menu"
	"server/internal/model/manage_role_menu"
	"server/internal/svc"
	menu_types "server/internal/types/manage/menu"
	types "server/internal/types/manage/role"
)

type SetMenus struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetMenus(ctx context.Context, svcCtx *svc.ServiceContext) *SetMenus {
	return &SetMenus{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetMenus) SetMenus(req *types.SetMenusRequest) (resp *types.SetMenusResponse, err error) {
	var datas []*manage_role_menu.ManageRoleMenu
	for _, v := range req.MenuIds {
		datas = append(datas, &manage_role_menu.ManageRoleMenu{
			RoleId:     int64(req.RoleId),
			MenuId:     int64(v),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
	}

	if len(datas) == 0 {
		return
	}

	if err = l.svcCtx.Model.ManageRoleMenu.DeleteByCondition(l.ctx, nil, condition.Condition{
		Field:    "role_id",
		Operator: condition.Equal,
		Value:    req.RoleId,
	}); err != nil {
		return
	}

	if err = l.svcCtx.Model.ManageRoleMenu.BulkInsert(l.ctx, nil, datas); err != nil {
		return nil, err
	}

	// update casbin_rule
	_, err = l.svcCtx.CasbinEnforcer.RemoveFilteredPolicy(0, cast.ToString(req.RoleId))
	if err != nil {
		return nil, errors.New("fail to remove filtered policy: " + err.Error())
	}
	// load policies
	err = l.svcCtx.CasbinEnforcer.LoadPolicy()
	if err != nil {
		return nil, errors.New("fail to load policy: " + err.Error())
	}

	// add casbin_rule
	var newPolicies [][]string
	// get menu perms
	menus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.New(condition.Condition{
		Field:    "id",
		Operator: condition.In,
		Value:    req.MenuIds,
	})...)
	if err != nil {
		return nil, err
	}
	for _, v := range menus {
		var permissions []menu_types.Permission
		menu.Unmarshal(v.Permissions.String, &permissions)
		for _, perm := range permissions {
			newPolicies = append(newPolicies, []string{cast.ToString(req.RoleId), perm.Code})
		}
	}

	var b bool
	if len(newPolicies) > 0 {
		b, _ = l.svcCtx.CasbinEnforcer.AddPolicies(newPolicies)
		if !b {
			return nil, errors.New("fail to add policies")
		}
		// load policies
		err = l.svcCtx.CasbinEnforcer.LoadPolicy()
	}
	return
}
