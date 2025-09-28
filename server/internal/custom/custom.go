package custom

import (
	"context"

	"github.com/jzero-io/jzero/core/stores/migrate"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/server/internal/config"
	"github.com/jzero-io/jzero-admin/server/internal/errcodes"
)

type Custom struct {
	Server *rest.Server
}

func New(server *rest.Server) *Custom {
	return &Custom{
		Server: server,
	}
}

// Init Please add custom logic here.
func (c *Custom) Init(cc configurator.Configurator[config.Config]) error {
	cfg, err := cc.GetConfig()
	if err != nil {
		return err
	}

	// register errcodes
	errcodes.Register()

	// migrate database
	if err := migrate.Migrate(context.Background(), cfg.Sqlx.SqlConf); err != nil {
		return err
	}

	c.AddRoutes(c.Server)
	return nil
}

// Start Please add custom logic here.
func (c *Custom) Start() {}

// Stop Please add shut down logic here.
func (c *Custom) Stop() {}
