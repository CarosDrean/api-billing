package bootstrap

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newEcho() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())

	return e
}
