package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"server/internal/config"
	"server/internal/custom"
	"server/internal/model"
	"server/pkg/localcache"
)

type ServiceContext struct {
	Config   config.Config
	SqlxConn sqlx.SqlConn
	Model    model.Model
	Cache    cache.Cache

	Custom *custom.Custom
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:   c,
		SqlxConn: MustSqlConn(c),

		Custom: custom.New(),
	}
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn)
	if c.CacheType == "local" {
		svcCtx.Cache = &localcache.Cache{
			Vals: make(map[string][]byte),
		}
	} else {
		// redis cache
	}
	return svcCtx
}
