package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type ProxyLogger struct{}

func NewProxyLogger() ProxyLogger {
	return ProxyLogger{}
}

func (l ProxyLogger) Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("request with url: %s has been redirected to gitea", c.Request().URL.String())
		return next(c)
	}
}
