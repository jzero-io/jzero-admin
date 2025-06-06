package auth

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/auth"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_menu"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role_menu"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_user_role"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/auth"
)

type GetUserInfo struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetUserInfo(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetUserInfo {
	return &GetUserInfo{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *GetUserInfo) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	info, err := auth.Info(l.ctx)
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.Model.ManageUser.FindOne(l.ctx, nil, uint64(info.Id))
	if err != nil {
		return nil, err
	}

	userRoles, err := l.svcCtx.Model.ManageUserRole.FindByCondition(l.ctx, nil, condition.NewChain().
		Equal(manage_user_role.UserId, user.Id).
		Build()...)
	if err != nil {
		return nil, err
	}

	var roleIds []int64
	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleId)
	}

	roles, err := l.svcCtx.Model.ManageRole.FindByCondition(l.ctx, nil, condition.NewChain().
		In(manage_role.Id, roleIds).
		Build()...)
	if err != nil {
		return nil, err
	}
	var roleCodes []string
	for _, role := range roles {
		roleCodes = append(roleCodes, role.Code)
	}

	// get role buttons
	roleMenus, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		In(manage_role_menu.RoleId, roleIds).
		Build()...)
	if err != nil {
		return nil, err
	}
	var menuIds []int64
	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.MenuId)
	}
	menus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
		In(manage_menu.Id, menuIds).
		Equal(manage_menu.Status, "1").
		Equal(manage_menu.MenuType, "3").
		Build()...)
	if err != nil {
		return nil, err
	}
	buttons := make([]string, 0)
	for _, menu := range menus {
		buttons = append(buttons, menu.ButtonCode.String)
	}

	return &types.GetUserInfoResponse{
		UserId:   cast.ToString(user.Id),
		Username: user.Username,
		Roles:    roleCodes,
		Buttons:  buttons,
	}, nil
}
