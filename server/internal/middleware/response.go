package middleware

import (
	"context"
)

type Body struct {
	Data any    `json:"data"`
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func ResponseMiddleware(_ context.Context, data any) any {
	return Body{
		Data: data,
		Code: "0000",
		Msg:  "success",
	}
}
