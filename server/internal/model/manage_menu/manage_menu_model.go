package manage_menu

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageMenuModel = (*customManageMenuModel)(nil)

type (
	// ManageMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageMenuModel.
	ManageMenuModel interface {
		manageMenuModel
	}

	customManageMenuModel struct {
		*defaultManageMenuModel
	}
)

// NewManageMenuModel returns a model for the database table.
func NewManageMenuModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageMenuModel {
	return &customManageMenuModel{
		defaultManageMenuModel: newManageMenuModel(conn, op...),
	}
}
