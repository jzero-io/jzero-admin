package manage_role

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageRoleModel = (*customManageRoleModel)(nil)

type (
	// ManageRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageRoleModel.
	ManageRoleModel interface {
		manageRoleModel
	}

	customManageRoleModel struct {
		*defaultManageRoleModel
	}
)

// NewManageRoleModel returns a model for the database table.
func NewManageRoleModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageRoleModel {
	return &customManageRoleModel{
		defaultManageRoleModel: newManageRoleModel(conn, op...),
	}
}
