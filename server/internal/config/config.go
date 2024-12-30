package config

import (
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/server/internal/i18n"
)

var C Config

type Config struct {
	Rest   RestConf
	Log    LogConf
	Banner BannerConf

	DatabaseType string `json:",default=mysql"`
	Mysql        MysqlConf
	Sqlite       SqliteConf

	CacheType string    `json:",default=local"`
	Redis     RedisConf `json:",optional"`

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
	Database string `json:",default=jzero"`
	Host     string `json:",default=127.0.0.1"`
	Port     int    `json:",default=3306"`
	Username string `json:",default=root"`
	Password string `json:",default=123456"`
}

type RedisConf struct {
	Host     string `json:","`
	Type     string `json:",default=node"`
	Pass     string `json:",optional,"`
	Tls      bool   `json:",optional"`
	NonBlock bool   `json:",default=true"`

	// PingTimeout is the timeout for ping redis.
	PingTimeout time.Duration `json:",default=1s"`
}

type Jwt struct {
	AccessSecret  string `json:",default=jzero-admin"`
	AccessExpire  int    `json:",default=7200"`
	RefreshExpire int    `json:",default=86400"`
}
