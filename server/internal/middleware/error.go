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

func ErrorMiddleware(ctx context.Context, err error) (int, any) {
	logx.Error("Request Error Context: %v", ctx)

	if errors.Is(err, jwt.ErrTokenExpired) {
		logx.Error(err)
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
