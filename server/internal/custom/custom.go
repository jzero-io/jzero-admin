package custom

import "github.com/jzero-io/jzero-admin/server/internal/errcodes"

type Custom struct {
}

func New() *Custom {
	return &Custom{}
}

// Start Please add custom logic here.
func (c *Custom) Start() {
	errcodes.Register()
}

// Stop Please add shut down logic here.
func (c *Custom) Stop() {}
