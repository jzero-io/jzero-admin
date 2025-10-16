package custom

type Custom struct{}

func New() *Custom {
	return &Custom{}
}

// Init Please add custom logic here.
func (c *Custom) Init() error {
	return nil
}

// Start Please add custom logic here.
func (c *Custom) Start() {}

// Stop Please add shut down logic here.
func (c *Custom) Stop() {}
