package svc

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/config"
)

func (sc *ServiceContext) GetConfig() (config.Config, error) {
	return sc.Config.GetConfig()
}

func (sc *ServiceContext) MustGetConfig() config.Config {
	c, err := sc.Config.GetConfig()
	logx.Must(err)
	return c
}

func (sc *ServiceContext) SetConfigListener(c config.Config, cc configurator.Configurator[config.Config]) {
	cc.AddListener(func() {
		logx.Infof("config file changed")
		if v, err := cc.GetConfig(); err == nil {
			if v.Log.Level != c.Log.Level {
				logx.Infof("log level changed: %s", v.Log.Level)
				switch v.Log.Level {
				case "debug":
					logx.SetLevel(logx.DebugLevel)
				case "info":
					logx.SetLevel(logx.InfoLevel)
				case "error":
					logx.SetLevel(logx.ErrorLevel)
				case "severe":
					logx.SetLevel(logx.SevereLevel)
				}
			}
		}
	})
}
