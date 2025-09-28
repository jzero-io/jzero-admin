package svc

import (
	"net/http"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/casbin/casbin/v2"
	"github.com/jzero-io/jzero/core/stores/cache"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/jzero-io/jzero/core/stores/redis"
	"github.com/pkg/errors"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/internal/config"
	"github.com/jzero-io/jzero-admin/server/internal/i18n"
	"github.com/jzero-io/jzero-admin/server/internal/model"
)

type ServiceContext struct {
	Config         configurator.Configurator[config.Config]
	SqlxConn       sqlx.SqlConn
	Model          model.Model
	Redis          *redis.Redis
	Cache          cache.Cache
	CasbinEnforcer *casbin.Enforcer
	Trans          *i18n.Translator
	Middleware
}

func NewServiceContext(cc configurator.Configurator[config.Config], route2Code func(r *http.Request) string) *ServiceContext {
	svcCtx := &ServiceContext{
		Config: cc,
	}
	svcCtx.SetConfigListener()
	svcCtx.SqlxConn = modelx.MustNewConn(svcCtx.MustGetConfig().Sqlx.SqlConf)
	if svcCtx.MustGetConfig().CacheType == "redis" {
		svcCtx.Redis = redis.MustNewRedis(svcCtx.MustGetConfig().Redis)
	} else {
		miniRedis, err := miniredis.Run()
		logx.Must(err)
		svcCtx.Redis = redis.MustNewRedis(redis.RedisConf{
			Type: redis.NodeType,
			Host: miniRedis.Addr(),
		})
	}
	svcCtx.Cache = cache.NewRedisNode(svcCtx.Redis, errors.New("cache not found"), cache.WithExpiry(time.Duration(5)*time.Second))
	svcCtx.CasbinEnforcer = MustCasbinEnforcer(svcCtx)
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn, modelx.WithCachedConn(modelx.NewConnWithCache(svcCtx.SqlxConn, svcCtx.Cache)))
	svcCtx.Trans = i18n.NewTranslator(svcCtx.MustGetConfig().I18n, i18n.LocaleFS)
	return svcCtx
}
