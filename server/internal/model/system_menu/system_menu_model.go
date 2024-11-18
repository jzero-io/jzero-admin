package system_menu

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemMenuModel = (*customSystemMenuModel)(nil)

type (
	// SystemMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemMenuModel.
	SystemMenuModel interface {
		systemMenuModel
	}

	customSystemMenuModel struct {
		*defaultSystemMenuModel
	}
)

// NewSystemMenuModel returns a model for the database table.
func NewSystemMenuModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) SystemMenuModel {
	return &customSystemMenuModel{
		defaultSystemMenuModel: newSystemMenuModel(conn, op...),
	}
}
