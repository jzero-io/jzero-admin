package serverless

var (
	PreProcessSqlFunc   func(version uint, content string) string
	BeforeMigrateUpFunc = map[uint]func(version uint) error{}
	AfterMigrateUpFunc  = map[uint]func(version uint) error{}
)
