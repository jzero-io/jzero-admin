package manage_email

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageEmailModel = (*customManageEmailModel)(nil)

type (
	// ManageEmailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageEmailModel.
	ManageEmailModel interface {
		manageEmailModel
	}

	customManageEmailModel struct {
		*defaultManageEmailModel
	}
)

// NewManageEmailModel returns a model for the database table.
func NewManageEmailModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageEmailModel {
	return &customManageEmailModel{
		defaultManageEmailModel: newManageEmailModel(conn, op...),
	}
}
