package system_role_menu

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemRoleMenuModel = (*customSystemRoleMenuModel)(nil)

type (
	// SystemRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemRoleMenuModel.
	SystemRoleMenuModel interface {
		systemRoleMenuModel
		WithSession(session sqlx.Session) SystemRoleMenuModel
	}

	customSystemRoleMenuModel struct {
		*defaultSystemRoleMenuModel
	}
)

// NewSystemRoleMenuModel returns a model for the database table.
func NewSystemRoleMenuModel(conn sqlx.SqlConn) SystemRoleMenuModel {
	return &customSystemRoleMenuModel{
		defaultSystemRoleMenuModel: newSystemRoleMenuModel(conn),
	}
}

func (m *customSystemRoleMenuModel) WithSession(session sqlx.Session) SystemRoleMenuModel {
	return NewSystemRoleMenuModel(sqlx.NewSqlConnFromSession(session))
}
