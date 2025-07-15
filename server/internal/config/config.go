package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/server/internal/i18n"
)

type Config struct {
	Rest      RestConf
	Jwt       Jwt
	Log       LogConf
	Banner    BannerConf
	Sqlx      SqlxConf
	CacheType string          `json:",default=local"`
	Redis     redis.RedisConf `json:",optional"`
	I18n      i18n.I18nConf   `json:",optional"`
}

type RestConf struct {
	rest.RestConf
}

type Jwt struct {
	AccessSecret  string `json:",default=jzero-admin"`
	AccessExpire  int    `json:",default=7200"`
	RefreshExpire int    `json:",default=86400"`
}

type LogConf struct {
	logx.LogConf
}

type BannerConf struct {
	Text     string `json:",default=JZERO"`
	Color    string `json:",default=green"`
	FontName string `json:",default=starwars,options=big|larry3d|starwars|standard"`
}

type SqlxConf struct {
	sqlx.SqlConf
}
