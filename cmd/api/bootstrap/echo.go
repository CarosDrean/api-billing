package bootstrap

import (
	"api-billing/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newEcho(conf model.Configuration, errorHandler echo.HTTPErrorHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.AllowedOrigins,
		AllowMethods: conf.AllowedMethods,
	}))

	e.HTTPErrorHandler = errorHandler

	return e
}
