package tables

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormConn() (*gorm.DB, error) {
	viper.SetConfigFile("etc/etc.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.AutomaticEnv()

	switch viper.GetString("databaseType") {
	case "mysql":
		return gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			viper.GetString("mysql.username"),
			viper.GetString("mysql.password"),
			viper.GetString("mysql.host")+":"+cast.ToString(viper.GetString("mysql.port")),
			viper.GetString("mysql.database"))), &gorm.Config{})
	case "sqlite":
		db, err := gorm.Open(sqlite.Open(viper.GetString("sqlite.path")), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	return nil, errors.Errorf("not support database [%s]", viper.GetString("databaseType"))
}

// Migrate https://gorm.io/docs/models.html
func Migrate(gormConn *gorm.DB) error {
	return gormConn.AutoMigrate(&CasbinRule{})
}
