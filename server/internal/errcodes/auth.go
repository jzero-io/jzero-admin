package errcodes

import "github.com/jzero-io/jzero/core/status"

const (
	RefreshTokenExpiredCode = 40102
)

func init() {
	status.Register(RefreshTokenExpiredCode)
}
