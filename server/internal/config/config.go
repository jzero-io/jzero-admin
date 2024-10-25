package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var C Config

type Config struct {
	Rest   RestConf
	Log    LogConf
	Banner BannerConf

	DatabaseType string `json:",default=mysql"`
	Mysql        MysqlConf
	Sqlite       SqliteConf

	Jwt Jwt

	CacheType string `json:",default=local"`
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

type Jwt struct {
	AccessSecret  string `json:",default=jzero-admin"`
	AccessExpire  int    `json:",default=7200"`
	RefreshExpire int    `json:",default=86400"`
}
