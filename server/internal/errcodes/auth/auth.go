package auth

import "github.com/jzero-io/jzero/core/status"

const (
	RefreshTokenExpiredCode = 40102
)

func RegisterAuth() {
	status.Register(RefreshTokenExpiredCode)
}
