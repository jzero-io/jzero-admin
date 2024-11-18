package svc

import (
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"

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
	if c.CacheType == "local" {
		svcCtx.Cache = &localcache.Cache{
			Vals: make(map[string][]byte),
		}
	} else {
		// redis cache
		singleFlights := syncx.NewSingleFlight()
		stats := cache.NewStat("redis-cache")
		svcCtx.Cache = cache.NewNode(&redis.Redis{
			Addr: svcCtx.Config.Redis.Host,
			Type: svcCtx.Config.Redis.Type,
			Pass: svcCtx.Config.Redis.Pass,
		}, singleFlights, stats, errors.New("no cache"))
	}
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn, modelx.WithCachedConn(sqlc.NewConnWithCache(svcCtx.SqlxConn, svcCtx.Cache)))
	return svcCtx
}
