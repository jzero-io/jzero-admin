package svc

import (
	"net/http"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/jzero-io/jzero-contrib/cache"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/pkg/errors"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	zerocache "github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/internal/config"
	"github.com/jzero-io/jzero-admin/server/internal/custom"
	"github.com/jzero-io/jzero-admin/server/internal/i18n"
	"github.com/jzero-io/jzero-admin/server/internal/model"
)

type ServiceContext struct {
	Config         configurator.Configurator[config.Config]
	SqlxConn       sqlx.SqlConn
	Model          model.Model
	Cache          cache.Cache
	CasbinEnforcer *casbin.Enforcer
	Trans          *i18n.Translator
	Middleware

	Custom *custom.Custom
}

func NewServiceContext(cc configurator.Configurator[config.Config], route2Code func(r *http.Request) string) *ServiceContext {
	svcCtx := &ServiceContext{
		Config: cc,
		Custom: custom.New(),
	}
	svcCtx.SetConfigListener()
	svcCtx.SqlxConn = MustSqlConn(svcCtx.MustGetConfig())
	if svcCtx.MustGetConfig().CacheType == "local" {
		svcCtx.Cache = cache.NewSyncMap(errors.New("cache not found"))
	} else {
		// redis cache
		rds := redis.MustNewRedis(redis.RedisConf(svcCtx.MustGetConfig().Redis))
		svcCtx.Cache = cache.NewRedisNode(rds, errors.New("cache not found"), zerocache.WithExpiry(time.Duration(5)*time.Second))
	}

	svcCtx.CasbinEnforcer = MustCasbinEnforcer(svcCtx)
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn, modelx.WithCachedConn(sqlc.NewConnWithCache(svcCtx.SqlxConn, svcCtx.Cache)))
	svcCtx.Middleware = NewMiddleware(svcCtx, route2Code)
	return svcCtx
}
