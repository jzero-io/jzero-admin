package golang

var (
	BeforeMigrateUpFunc = map[uint]func(version uint) error{}
	AfterMigrateUpFunc  = map[uint]func(version uint) error{}
)
