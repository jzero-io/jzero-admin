package system_user

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemUserModel = (*customSystemUserModel)(nil)

type (
	// SystemUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemUserModel.
	SystemUserModel interface {
		systemUserModel
	}

	customSystemUserModel struct {
		*defaultSystemUserModel
	}
)

// NewSystemUserModel returns a model for the database table.
func NewSystemUserModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) SystemUserModel {
	return &customSystemUserModel{
		defaultSystemUserModel: newSystemUserModel(conn, op...),
	}
}
