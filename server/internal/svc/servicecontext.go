package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"server/internal/config"
	"server/internal/custom"
)

type ServiceContext struct {
	Config   config.Config
	SqlxConn sqlx.SqlConn

	Custom *custom.Custom
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		SqlxConn: MustSqlConn(c),

		Custom: custom.New(),
	}
}
