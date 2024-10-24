package middleware

import (
	"context"
	"net/http"

	"github.com/spf13/cast"
)

type Body struct {
	Data any    `json:"data"`
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func ResponseMiddleware(_ context.Context, data any) any {
	return Body{
		Data: data,
		Code: cast.ToString(http.StatusOK),
		Msg:  "success",
	}
}
