package system_role

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemRoleModel = (*customSystemRoleModel)(nil)

type (
	// SystemRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemRoleModel.
	SystemRoleModel interface {
		systemRoleModel
	}

	customSystemRoleModel struct {
		*defaultSystemRoleModel
	}
)

// NewSystemRoleModel returns a model for the database table.
func NewSystemRoleModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) SystemRoleModel {
	return &customSystemRoleModel{
		defaultSystemRoleModel: newSystemRoleModel(conn, op...),
	}
}
