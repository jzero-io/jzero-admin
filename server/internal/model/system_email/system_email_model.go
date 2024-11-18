package system_email

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemEmailModel = (*customSystemEmailModel)(nil)

type (
	// SystemEmailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemEmailModel.
	SystemEmailModel interface {
		systemEmailModel
	}

	customSystemEmailModel struct {
		*defaultSystemEmailModel
	}
)

// NewSystemEmailModel returns a model for the database table.
func NewSystemEmailModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) SystemEmailModel {
	return &customSystemEmailModel{
		defaultSystemEmailModel: newSystemEmailModel(conn, op...),
	}
}
