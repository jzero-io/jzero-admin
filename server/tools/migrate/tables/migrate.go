package tables

import (
	"fmt"
	"log"
	"os"

	"github.com/a8m/envsubst"
	"github.com/glebarez/sqlite"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/conf"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DatabaseType string `json:",default=mysql"`
	Sqlite       SqliteConf
	Mysql        MysqlConf
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

func NewGormConn() (*gorm.DB, error) {
	var c Config
	conf.MustLoad("etc/etc.yaml", &c, conf.UseEnv())

	switch c.DatabaseType {
	case "mysql":
		return gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Mysql.Username,
			c.Mysql.Password,
			c.Mysql.Host+":"+cast.ToString(c.Mysql.Port),
			c.Mysql.Database)), &gorm.Config{})
	case "sqlite":
		db, err := gorm.Open(sqlite.Open(c.Sqlite.Path), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	return nil, errors.Errorf("not support database [%s]", c.DatabaseType)
}

// Migrate https://gorm.io/docs/models.html
func Migrate(gormConn *gorm.DB) error {
	return gormConn.AutoMigrate(
		&CasbinRule{},
		&SystemUser{},
		&SystemRole{},
		&SystemUserRole{},
		&SystemMenu{},
		&SystemRoleMenu{},
	)
}

func init() {
	data, err := envsubst.ReadFile("etc/.env.yaml")
	if err != nil {
		log.Fatalf("envsubst error: %v", err)
	}
	var env map[string]any
	err = yaml.Unmarshal(data, &env)
	if err != nil {
		log.Fatalf("yaml unmarshal error: %v", err)
	}

	for k, v := range env {
		_ = os.Setenv(k, cast.ToString(v))
	}
}
