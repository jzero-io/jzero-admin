package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"server/internal/config"
	"server/internal/custom"
	"server/internal/model"
)

type ServiceContext struct {
	Config   config.Config
	SqlxConn sqlx.SqlConn
	Model    model.Model

	Custom *custom.Custom
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:   c,
		SqlxConn: MustSqlConn(c),

		Custom: custom.New(),
	}
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn)
	return svcCtx
}
