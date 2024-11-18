package system_role_menu

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemRoleMenuModel = (*customSystemRoleMenuModel)(nil)

type (
	// SystemRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemRoleMenuModel.
	SystemRoleMenuModel interface {
		systemRoleMenuModel
	}

	customSystemRoleMenuModel struct {
		*defaultSystemRoleMenuModel
	}
)

// NewSystemRoleMenuModel returns a model for the database table.
func NewSystemRoleMenuModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) SystemRoleMenuModel {
	return &customSystemRoleMenuModel{
		defaultSystemRoleMenuModel: newSystemRoleMenuModel(conn, op...),
	}
}
