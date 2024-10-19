package svc

import (
	"fmt"

	_ "github.com/glebarez/sqlite"
	"github.com/huandu/go-sqlbuilder"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

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
