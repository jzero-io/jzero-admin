package system_user_role

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemUserRoleModel = (*customSystemUserRoleModel)(nil)

type (
	// SystemUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemUserRoleModel.
	SystemUserRoleModel interface {
		systemUserRoleModel
	}

	customSystemUserRoleModel struct {
		*defaultSystemUserRoleModel
	}
)

// NewSystemUserRoleModel returns a model for the database table.
func NewSystemUserRoleModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) SystemUserRoleModel {
	return &customSystemUserRoleModel{
		defaultSystemUserRoleModel: newSystemUserRoleModel(conn, op...),
	}
}
