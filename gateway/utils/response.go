package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	SuccessStatus = "success"
	FailStatus    = "fail"
	ErrorStatus   = "error"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func HandleGrpcStatus(ctx echo.Context, err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return Response(ctx, http.StatusInternalServerError, "Internal error", nil)
	}

	code := http.StatusInternalServerError
	switch st.Code() {
	case codes.InvalidArgument:
		code = http.StatusBadRequest
	case codes.NotFound:
		code = http.StatusNotFound
	case codes.AlreadyExists:
		code = http.StatusConflict
	}

	return Response(ctx, code, st.Message(), nil)
}

func Response(ctx echo.Context, code int, message string, object any) error {
	status := SuccessStatus
	switch {
	case code >= 400 && code < 500:
		status = FailStatus
	case code >= 500:
		status = ErrorStatus
	}

	return ctx.JSON(code, response{
		Status:  status,
		Data:    object,
		Message: message,
	})
}
