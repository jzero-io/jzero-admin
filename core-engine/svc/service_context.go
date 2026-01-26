package svc

import (
	"net/http"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/casbin/casbin/v2"
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero/core/stores/cache"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/core-engine/config"
	"github.com/jzero-io/jzero-admin/core-engine/i18n"
)

type ServiceContext struct {
	Config         config.Config
	SqlxConn       sqlx.SqlConn
	Redis          *redis.Redis
	Cache          cache.Cache
	CasbinEnforcer *casbin.Enforcer
	Trans          *i18n.Translator
	Middleware
}

type ServiceContextOpts struct {
	Serverless bool
}

func WithServerless(serverless bool) opts.Opt[ServiceContextOpts] {
	return func(opts *ServiceContextOpts) {
		opts.Serverless = serverless
	}
}

func (opts ServiceContextOpts) DefaultOptions() ServiceContextOpts {
	return ServiceContextOpts{}
}

func NewServiceContext(c config.Config, route2code func(r *http.Request) string, op ...opts.Opt[ServiceContextOpts]) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:   c,
		SqlxConn: modelx.MustNewConn(c.Sqlx.SqlConf),
	}

	if c.Redis.MiniRedis {
		miniRedis, err := miniredis.Run()
		logx.Must(err)
		svcCtx.Redis = redis.MustNewRedis(redis.RedisConf{
			Type: redis.NodeType,
			Host: miniRedis.Addr(),
		})
	} else if c.Redis.Host != "" {
		svcCtx.Redis = redis.MustNewRedis(c.Redis.RedisConf)
	}

	svcCtx.Cache = cache.NewRedisNode(svcCtx.Redis, errors.New("cache not found"), cache.WithExpiry(time.Duration(5)*time.Second))
	svcCtx.CasbinEnforcer = MustCasbinEnforcer(svcCtx)
	svcCtx.Trans = i18n.NewTranslator(c.I18n, i18n.LocaleFS)
	svcCtx.Middleware = NewMiddleware(svcCtx, route2code)
	return svcCtx
}
