package manage_role_menu

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageRoleMenuModel = (*customManageRoleMenuModel)(nil)

type (
	// ManageRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageRoleMenuModel.
	ManageRoleMenuModel interface {
		manageRoleMenuModel
	}

	customManageRoleMenuModel struct {
		*defaultManageRoleMenuModel
	}
)

// NewManageRoleMenuModel returns a model for the database table.
func NewManageRoleMenuModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageRoleMenuModel {
	return &customManageRoleMenuModel{
		defaultManageRoleMenuModel: newManageRoleMenuModel(conn, op...),
	}
}
