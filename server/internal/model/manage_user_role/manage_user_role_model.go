package manage_user_role

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageUserRoleModel = (*customManageUserRoleModel)(nil)

type (
	// ManageUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageUserRoleModel.
	ManageUserRoleModel interface {
		manageUserRoleModel
	}

	customManageUserRoleModel struct {
		*defaultManageUserRoleModel
	}
)

// NewManageUserRoleModel returns a model for the database table.
func NewManageUserRoleModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageUserRoleModel {
	return &customManageUserRoleModel{
		defaultManageUserRoleModel: newManageUserRoleModel(conn, op...),
	}
}
