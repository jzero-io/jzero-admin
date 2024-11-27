package svc

import (
	"net/http"

	"github.com/jzero-io/jzero-contrib/cache"
	"github.com/jzero-io/jzero-contrib/cache/sync"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/pkg/errors"
	zerocache "github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/gorm"

	"server/internal/config"
	"server/internal/custom"
	"server/internal/middleware"
	"server/internal/model"
)

type ServiceContext struct {
	Config   config.Config
	SqlxConn sqlx.SqlConn
	GormDB   *gorm.DB // TODO: 去掉 gorm
	Model    model.Model
	Cache    cache.Cache
	Custom   *custom.Custom
	middleware.Middleware
}

func NewServiceContext(c config.Config, route2Code func(r *http.Request) string) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:   c,
		SqlxConn: MustSqlConn(c),
		GormDB:   MustGormDB(c),
		Custom:   custom.New(),
	}
	if c.CacheType == "local" {
		svcCtx.Cache = sync.New(errors.New("cache not found"))
	} else {
		// redis cache
		singleFlights := syncx.NewSingleFlight()
		stats := zerocache.NewStat("redis-cache")
		svcCtx.Cache = zerocache.NewNode(&redis.Redis{
			Addr: svcCtx.Config.Redis.Host,
			Type: svcCtx.Config.Redis.Type,
			Pass: svcCtx.Config.Redis.Pass,
		}, singleFlights, stats, errors.New("no cache"))
	}
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn, modelx.WithCachedConn(sqlc.NewConnWithCache(svcCtx.SqlxConn, svcCtx.Cache)))
	svcCtx.Middleware = middleware.NewMiddleware(svcCtx.Cache, svcCtx.SqlxConn, svcCtx.GormDB, route2Code)
	return svcCtx
}
