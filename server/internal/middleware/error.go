package middleware

import (
	"net/http"

	"github.com/jzero-io/jzero-contrib/status"
	"github.com/spf13/cast"
)

func ErrorMiddleware(err error) (int, any) {
	fromError := status.FromError(err)
	return http.StatusOK, Body{
		Data: nil,
		Code: cast.ToString(int(fromError.Code())),
		Msg:  fromError.Error(),
	}
}
