package custom

import (
	"{{.Module}}/internal/config"
)

type Custom struct{
    c config.Config
	// Please add custom fields here.
}

func New(c config.Config) *Custom {
	return &Custom{
		c: c,
	}
}

// Init Please add custom logic here.
func (c *Custom) Init() error {
	return nil
}

// Start Please add custom logic here.
func (c *Custom) Start() {}

// Stop Please add shut down logic here.
func (c *Custom) Stop() {}
