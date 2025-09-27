package custom

import (
	"context"

	"github.com/jzero-io/jzero/core/stores/migrate"
	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/server/internal/errcodes"
	"github.com/jzero-io/jzero-admin/server/internal/global"
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
func (c *Custom) Init() error {
	errcodes.Register()
	if err := migrate.Migrate(context.Background(), global.ServiceContext.MustGetConfig().Sqlx.SqlConf); err != nil {
		return err
	}

	c.AddRoutes(c.Server)
	return nil
}

// Start Please add custom logic here.
func (c *Custom) Start() {}

// Stop Please add shut down logic here.
func (c *Custom) Stop() {}
