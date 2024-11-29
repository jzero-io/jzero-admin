package manage

import "github.com/jzero-io/jzero-contrib/status"

type ErrorPair struct {
	Code    status.Code
	Message string
}

const (
	ExistSubMenuCode = 10311
)

var (
	ExistSubMenuError = ErrorPair{Code: ExistSubMenuCode, Message: "Exist submenu error"}
)

func RegisterManage() {
	status.RegisterWithMessage(ExistSubMenuError.Code, ExistSubMenuError.Message)
}
