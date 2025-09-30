package custom

import (
	"github.com/jzero-io/jzero-admin/server/internal/config"
	"github.com/jzero-io/jzero-admin/server/internal/errcodes"
)

type Custom struct {
	c config.Config
}

func New(c config.Config) *Custom {
	return &Custom{
		c: c,
	}
}

// Init Please add custom logic here.
func (c *Custom) Init() error {
	// register errcodes
	errcodes.Register()
	return nil
}

// Start Please add custom logic here.
func (c *Custom) Start() {}

// Stop Please add shut down logic here.
func (c *Custom) Stop() {}
