package system_role

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemRoleModel = (*customSystemRoleModel)(nil)

type (
	// SystemRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemRoleModel.
	SystemRoleModel interface {
		systemRoleModel
		WithSession(session sqlx.Session) SystemRoleModel
	}

	customSystemRoleModel struct {
		*defaultSystemRoleModel
	}
)

// NewSystemRoleModel returns a model for the database table.
func NewSystemRoleModel(conn sqlx.SqlConn) SystemRoleModel {
	return &customSystemRoleModel{
		defaultSystemRoleModel: newSystemRoleModel(conn),
	}
}

func (m *customSystemRoleModel) WithSession(session sqlx.Session) SystemRoleModel {
	return NewSystemRoleModel(sqlx.NewSqlConnFromSession(session))
}
