package custom

import (
	"context"
	"os"

	"github.com/jzero-io/jzero/core/stores/migrate"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/config"
	"github.com/jzero-io/jzero-admin/server/internal/errcodes"
)

type Custom struct{}

func New() *Custom {
	return &Custom{}
}

// Init Please add custom logic here.
func (c *Custom) Init(cfg config.Config) error {
	// register errcodes
	errcodes.Register()

	// migrate database
	if err := migrate.MigrateUp(context.Background(), cfg.Sqlx.SqlConf, migrate.WithSource(func() string {
		switch cfg.Sqlx.DriverName {
		case "mysql":
			return "file://desc/sql_migration"
		case "pgx":
			return "file://desc/sql_migration/postgresql"
		default:
			return "file://desc/sql_migration"
		}
	}())); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logx.Infof("migration source not exist, skip migration")
			return nil
		}
		return err
	}

	return nil
}

// Start Please add custom logic here.
func (c *Custom) Start() {}

// Stop Please add shut down logic here.
func (c *Custom) Stop() {}
