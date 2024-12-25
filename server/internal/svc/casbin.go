package svc

import (
	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/zeromicro/go-zero/core/logx"
)

const CasbinModelConf = `[request_definition]
r = sub, obj

[policy_definition]
p = sub, obj

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj`

func MustCasbinEnforcer(svcCtx *ServiceContext) *casbin.Enforcer {
	config, err := svcCtx.Config.GetConfig()
	logx.Must(err)

	db, err := svcCtx.SqlxConn.RawDB()
	logx.Must(err)

	adapter, err := sqladapter.NewAdapter(db, config.DatabaseType, "casbin_rule")
	logx.Must(err)

	casbinModel, err := casbinmodel.NewModelFromString(CasbinModelConf)
	logx.Must(err)

	casbinEnforcer, err := casbin.NewEnforcer(casbinModel, adapter)
	logx.Must(err)

	err = casbinEnforcer.LoadPolicy()
	logx.Must(err)

	return casbinEnforcer
}
