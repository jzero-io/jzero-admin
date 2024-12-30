package svc

import (
	"fmt"

	_ "github.com/glebarez/sqlite"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/internal/config"
)

func BuildDataSource(c config.Config) string {
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
	sqlConn := sqlx.NewSqlConn(c.DatabaseType, BuildDataSource(c))
	db, err := sqlConn.RawDB()
	logx.Must(err)
	logx.Must(db.Ping())
	return sqlConn
}
