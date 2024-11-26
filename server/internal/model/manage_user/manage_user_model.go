package manage_user

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageUserModel = (*customManageUserModel)(nil)

type (
	// ManageUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageUserModel.
	ManageUserModel interface {
		manageUserModel
	}

	customManageUserModel struct {
		*defaultManageUserModel
	}
)

// NewManageUserModel returns a model for the database table.
func NewManageUserModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageUserModel {
	return &customManageUserModel{
		defaultManageUserModel: newManageUserModel(conn, op...),
	}
}
