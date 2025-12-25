package errcodes

import "github.com/jzero-io/jzero/core/status"

const (
	ExistSubMenuCode = 10311
)

func init() {
	status.Register(ExistSubMenuCode)
}
