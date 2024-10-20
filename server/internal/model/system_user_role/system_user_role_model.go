package system_user_role

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemUserRoleModel = (*customSystemUserRoleModel)(nil)

type (
	// SystemUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemUserRoleModel.
	SystemUserRoleModel interface {
		systemUserRoleModel
		WithSession(session sqlx.Session) SystemUserRoleModel
	}

	customSystemUserRoleModel struct {
		*defaultSystemUserRoleModel
	}
)

// NewSystemUserRoleModel returns a model for the database table.
func NewSystemUserRoleModel(conn sqlx.SqlConn) SystemUserRoleModel {
	return &customSystemUserRoleModel{
		defaultSystemUserRoleModel: newSystemUserRoleModel(conn),
	}
}

func (m *customSystemUserRoleModel) WithSession(session sqlx.Session) SystemUserRoleModel {
	return NewSystemUserRoleModel(sqlx.NewSqlConnFromSession(session))
}
