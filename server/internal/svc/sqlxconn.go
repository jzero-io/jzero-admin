package svc

import (
	"github.com/huandu/go-sqlbuilder"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func MustSqlxConn(c config.SqlxConf) sqlx.SqlConn {
	sqlConn := sqlx.MustNewConn(c.SqlConf)
	db, err := sqlConn.RawDB()
	logx.Must(err)
	err = db.Ping()
	logx.Must(err)

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)

	setSqlbuilderFlavor(c.DriverName)
	return sqlConn
}

func setSqlbuilderFlavor(driverName string) {
	switch driverName {
	case "mysql":
		sqlbuilder.DefaultFlavor = sqlbuilder.MySQL
	case "pgx", "postgres":
		sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL
	}
}
