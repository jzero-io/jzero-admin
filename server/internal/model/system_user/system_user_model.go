package system_user

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemUserModel = (*customSystemUserModel)(nil)

type (
	// SystemUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemUserModel.
	SystemUserModel interface {
		systemUserModel
		WithSession(session sqlx.Session) SystemUserModel
	}

	customSystemUserModel struct {
		*defaultSystemUserModel
	}
)

// NewSystemUserModel returns a model for the database table.
func NewSystemUserModel(conn sqlx.SqlConn) SystemUserModel {
	return &customSystemUserModel{
		defaultSystemUserModel: newSystemUserModel(conn),
	}
}

func (m *customSystemUserModel) WithSession(session sqlx.Session) SystemUserModel {
	return NewSystemUserModel(sqlx.NewSqlConnFromSession(session))
}
