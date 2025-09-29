package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/jzero-io/jzero/core/stores/cache"
	"github.com/jzero-io/jzero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	SqlxConn       sqlx.SqlConn
	Redis          *redis.Redis
	Cache          cache.Cache
	CasbinEnforcer *casbin.Enforcer
}
