package custom

import (
	"context"

	"github.com/jzero-io/jzero-admin/core-engine/helper/migrate"
	"github.com/jzero-io/jzero-admin/server/internal/errcodes"
	"github.com/jzero-io/jzero-admin/server/internal/global"
)

type Custom struct {
}

func New() *Custom {
	return &Custom{}
}

// Init Please add custom logic here.
func (c *Custom) Init() error {
	if err := migrate.MigrateUp(context.Background(), global.ServiceContext.ConfigCenter.MustGetConfig().Sqlx.SqlConf); err != nil {
		return err
	}

	// register errcodes
	errcodes.Register()
	return nil
}

// Start Please add custom logic here.
func (c *Custom) Start() {}

// Stop Please add shut down logic here.
func (c *Custom) Stop() {}
