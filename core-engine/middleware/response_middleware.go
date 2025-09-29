package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jzero-io/jzero/core/status"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

type Body struct {
	Data any    `json:"data"`
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type OkMiddleware struct{}

func NewOkMiddleware() *OkMiddleware {
	return &OkMiddleware{}
}

type ErrorMiddleware struct{}

func NewErrorMiddleware() *ErrorMiddleware {
	return &ErrorMiddleware{}
}

func (e *ErrorMiddleware) Handle(ctx context.Context, err error) (int, any) {
	logx.WithContext(ctx).Errorf("request error: %v", err)

	if errors.Is(err, jwt.ErrTokenExpired) {
		return http.StatusOK, Body{
			Data: nil,
			Code: "40101",
			Msg:  "token expired to refresh token",
		}
	}

	fromError := status.FromError(err)
	return http.StatusOK, Body{
		Data: nil,
		Code: cast.ToString(int(fromError.Code())),
		Msg:  fromError.Error(),
	}
}

func (o *OkMiddleware) Handle(_ context.Context, data any) any {
	return Body{
		Data: data,
		Code: cast.ToString(http.StatusOK),
		Msg:  "success",
	}
}
