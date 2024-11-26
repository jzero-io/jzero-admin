package manage_menu_role_permission

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageMenuRolePermissionModel = (*customManageMenuRolePermissionModel)(nil)

type (
	// ManageMenuRolePermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageMenuRolePermissionModel.
	ManageMenuRolePermissionModel interface {
		manageMenuRolePermissionModel
	}

	customManageMenuRolePermissionModel struct {
		*defaultManageMenuRolePermissionModel
	}
)

// NewManageMenuRolePermissionModel returns a model for the database table.
func NewManageMenuRolePermissionModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageMenuRolePermissionModel {
	return &customManageMenuRolePermissionModel{
		defaultManageMenuRolePermissionModel: newManageMenuRolePermissionModel(conn, op...),
	}
}
