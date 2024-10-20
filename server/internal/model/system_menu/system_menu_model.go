package system_menu

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemMenuModel = (*customSystemMenuModel)(nil)

type (
	// SystemMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemMenuModel.
	SystemMenuModel interface {
		systemMenuModel
		WithSession(session sqlx.Session) SystemMenuModel
	}

	customSystemMenuModel struct {
		*defaultSystemMenuModel
	}
)

// NewSystemMenuModel returns a model for the database table.
func NewSystemMenuModel(conn sqlx.SqlConn) SystemMenuModel {
	return &customSystemMenuModel{
		defaultSystemMenuModel: newSystemMenuModel(conn),
	}
}

func (m *customSystemMenuModel) WithSession(session sqlx.Session) SystemMenuModel {
	return NewSystemMenuModel(sqlx.NewSqlConnFromSession(session))
}
