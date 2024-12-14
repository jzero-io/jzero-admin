package svc

import (
	"net/http"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/jzero-io/jzero-contrib/cache"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/pkg/errors"
	zerocache "github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"server/internal/config"
	"server/internal/custom"
	"server/internal/i18n"
	"server/internal/model"
)

type ServiceContext struct {
	Config         config.Config
	SqlxConn       sqlx.SqlConn
	Model          model.Model
	Cache          cache.Cache
	CasbinEnforcer *casbin.Enforcer
	Trans          *i18n.Translator
	Middleware

	Custom *custom.Custom
}

func NewServiceContext(c config.Config, route2Code func(r *http.Request) string) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:   c,
		SqlxConn: MustSqlConn(c),
		Custom:   custom.New(),
		Trans:    i18n.NewTranslator(c.I18n, i18n.LocaleFS),
	}
	if c.CacheType == "local" {
		svcCtx.Cache = cache.NewSyncMap(errors.New("cache not found"))
	} else {
		// redis cache
		rds := redis.MustNewRedis(redis.RedisConf(svcCtx.Config.Redis))
		svcCtx.Cache = cache.NewRedisNode(rds, errors.New("cache not found"), zerocache.WithExpiry(time.Duration(5)*time.Second))
	}

	svcCtx.CasbinEnforcer = MustCasbinEnforcer(svcCtx)
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn, modelx.WithCachedConn(sqlc.NewConnWithCache(svcCtx.SqlxConn, svcCtx.Cache)))
	svcCtx.Middleware = NewMiddleware(svcCtx, route2Code)
	return svcCtx
}
