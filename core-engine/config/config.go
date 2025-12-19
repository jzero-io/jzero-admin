package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/core-engine/i18n"
)

type Config struct {
	Rest   RestConf
	Jwt    JwtConf
	Log    LogConf
	Banner BannerConf
	Sqlx   SqlxConf
	Redis  RedisConf     `json:",optional"`
	I18n   i18n.I18nConf `json:",optional"`
}

type RestConf struct {
	rest.RestConf
}

type JwtConf struct {
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

type RedisConf struct {
	// MiniRedis only for testing
	MiniRedis bool `json:",default=false"`

	redis.RedisConf
}
