package system_email

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemEmailModel = (*customSystemEmailModel)(nil)

type (
	// SystemEmailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemEmailModel.
	SystemEmailModel interface {
		systemEmailModel
		WithSession(session sqlx.Session) SystemEmailModel
	}

	customSystemEmailModel struct {
		*defaultSystemEmailModel
	}
)

// NewSystemEmailModel returns a model for the database table.
func NewSystemEmailModel(conn sqlx.SqlConn) SystemEmailModel {
	return &customSystemEmailModel{
		defaultSystemEmailModel: newSystemEmailModel(conn),
	}
}

func (m *customSystemEmailModel) WithSession(session sqlx.Session) SystemEmailModel {
	return NewSystemEmailModel(sqlx.NewSqlConnFromSession(session))
}
