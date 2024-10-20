package casbin_rule

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CasbinRuleModel = (*customCasbinRuleModel)(nil)

type (
	// CasbinRuleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCasbinRuleModel.
	CasbinRuleModel interface {
		casbinRuleModel
		WithSession(session sqlx.Session) CasbinRuleModel
	}

	customCasbinRuleModel struct {
		*defaultCasbinRuleModel
	}
)

// NewCasbinRuleModel returns a model for the database table.
func NewCasbinRuleModel(conn sqlx.SqlConn) CasbinRuleModel {
	return &customCasbinRuleModel{
		defaultCasbinRuleModel: newCasbinRuleModel(conn),
	}
}

func (m *customCasbinRuleModel) WithSession(session sqlx.Session) CasbinRuleModel {
	return NewCasbinRuleModel(sqlx.NewSqlConnFromSession(session))
}
