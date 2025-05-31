package manage

import "github.com/jzero-io/jzero/core/status"

const (
	ExistSubMenuCode = 10311
)

func RegisterManage() {
	status.Register(ExistSubMenuCode)
}
