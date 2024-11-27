package middleware

import (
	"net/http"

	casbin "github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/jzero-io/jzero-contrib/cache"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
)

type Middleware struct {
	Authx rest.Middleware
}

func NewMiddleware(cache cache.Cache, sqlxConn sqlx.SqlConn, gormDB *gorm.DB, route2Code func(r *http.Request) string) Middleware {
	gormadapter.TurnOffAutoMigrate(gormDB)
	adapter, err := gormadapter.NewAdapterByDB(gormDB)
	if err != nil {
		logx.Must(err)
	}

	casbinModel, err := casbinmodel.NewModelFromString(CasbinModelConf)
	if err != nil {
		logx.Must(err)
	}

	casbinEnforcer, err := casbin.NewEnforcer(casbinModel, adapter)
	if err != nil {
		logx.Must(err)
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		logx.Must(err)
	}

	return Middleware{
		Authx: NewAuthxMiddleware(cache, sqlxConn, casbinEnforcer, route2Code).Handle,
	}
}

func Register(server *rest.Server) {
	httpx.SetOkHandler(ResponseMiddleware)
	httpx.SetErrorHandler(ErrorMiddleware)
	httpx.SetValidator(NewValidator())
}
