package svc

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/huandu/go-sqlbuilder"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/internal/config"
)

func BuildDataSource(c config.Config) string {
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
	case "postgres":
		sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL
		return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			c.Postgres.Username,
			c.Postgres.Password,
			c.Postgres.Host+":"+cast.ToString(c.Postgres.Port),
			c.Postgres.Database)
	}
	return ""
}

func MustSqlConn(c config.Config) sqlx.SqlConn {
	var sqlConn sqlx.SqlConn

	switch c.DatabaseType {
	case "mysql":
		sqlConn = sqlx.NewMysql(BuildDataSource(c))
	case "postgres":
		sqlConn = postgres.New(BuildDataSource(c))
	case "sqlite":
		sqlConn = sqlx.NewSqlConn(sqlite.DriverName, BuildDataSource(c))
	default:
		panic(fmt.Sprintf("not supported database type: %s", c.DatabaseType))
	}

	db, err := sqlConn.RawDB()
	logx.Must(err)
	logx.Must(db.Ping())
	return sqlConn
}
