package svc

import (
	"fmt"

	"github.com/glebarez/sqlite"
	_ "github.com/glebarez/sqlite"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"server/internal/config"
)

func buildDataSource(c config.Config) string {
	// set default sqlbuilder flavor and data source
	switch c.DatabaseType {
	case "mysql":
		sqlbuilder.DefaultFlavor = sqlbuilder.MySQL
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Mysql.Username,
			c.Mysql.Password,
			c.Mysql.Host+":"+cast.ToString(c.Mysql.Port),
			c.Mysql.Database)
	case "sqlite":
		sqlbuilder.DefaultFlavor = sqlbuilder.SQLite
		return c.Sqlite.Path
	}
	return ""
}

func MustSqlConn(c config.Config) sqlx.SqlConn {
	sqlConn := sqlx.NewSqlConn(c.DatabaseType, buildDataSource(c))
	_, err := sqlConn.Exec("select 1")
	if err != nil {
		panic(err)
	}
	return sqlConn
}

func MustGormDB(c config.Config) *gorm.DB {
	conn, err := newGormDB(c)
	if err != nil {
		panic(err)
	}
	return conn
}

func newGormDB(c config.Config) (*gorm.DB, error) {
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
