package manage_menu_permission

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageMenuPermissionModel = (*customManageMenuPermissionModel)(nil)

type (
	// ManageMenuPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageMenuPermissionModel.
	ManageMenuPermissionModel interface {
		manageMenuPermissionModel
	}

	customManageMenuPermissionModel struct {
		*defaultManageMenuPermissionModel
	}
)

// NewManageMenuPermissionModel returns a model for the database table.
func NewManageMenuPermissionModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageMenuPermissionModel {
	return &customManageMenuPermissionModel{
		defaultManageMenuPermissionModel: newManageMenuPermissionModel(conn, op...),
	}
}
