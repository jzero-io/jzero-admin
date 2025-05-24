package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/server/internal/i18n"
)

type Config struct {
	Rest   RestConf
	Log    LogConf
	Banner BannerConf

	DatabaseType string `json:",default=mysql"`
	Postgres     PostgresConf
	Mysql        MysqlConf
	Sqlite       SqliteConf

	CacheType string          `json:",default=local"`
	Redis     redis.RedisConf `json:",optional"`

	I18n i18n.Conf `json:",optional"`

	Jwt Jwt
}

type RestConf struct {
	rest.RestConf
}

type LogConf struct {
	logx.LogConf
}

type BannerConf struct {
	Text     string `json:",default=JZERO"`
	Color    string `json:",default=green"`
	FontName string `json:",default=starwars,options=big|larry3d|starwars|standard"`
}

type SqliteConf struct {
	Path string `json:",default=data/jzero.db"`
}

type MysqlConf struct {
	Database string `json:",default=jzeroadmin"`
	Host     string `json:",default=127.0.0.1"`
	Port     int    `json:",default=3306"`
	Username string `json:",default=root"`
	Password string `json:",default=123456"`
}

type PostgresConf struct {
	Database string `json:",default=jzeroadmin"`
	Host     string `json:",default=127.0.0.1"`
	Port     int    `json:",default=5432"`
	Username string `json:",default=root"`
	Password string `json:",default=123456"`
}

type Jwt struct {
	AccessSecret  string `json:",default=jzero-admin"`
	AccessExpire  int    `json:",default=7200"`
	RefreshExpire int    `json:",default=86400"`
}
