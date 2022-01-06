package invoice

import (
	"database/sql"

	"api-billing/domain/invoice"
	invoiceStorage "api-billing/insfraestructure/postgres/invoice"

	"github.com/labstack/echo/v4"
)

const (
	privateRoutePrefix = "v1/invoice"
)

func NewRoutes(app *echo.Echo, db *sql.DB) {
	useCase := invoice.New(invoiceStorage.New(db))

	handler := newHandler(useCase)

	privateRoutes(app, handler)
}

func privateRoutes(app *echo.Echo, handler handler) {
	api := app.Group(privateRoutePrefix)

	api.POST("", handler.create)
	api.GET("", handler.getAllByRangeCreatedAt)
	api.GET("/simulate", handler.GetPriceByMedicinesIDsAndSaleDate)
}
