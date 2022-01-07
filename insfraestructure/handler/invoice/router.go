package invoice

import (
	"database/sql"

	"api-billing/domain/invoice"
	"api-billing/domain/medicine"
	"api-billing/domain/promotion"
	invoiceStorage "api-billing/insfraestructure/postgres/invoice"
	medicineStorage "api-billing/insfraestructure/postgres/medicine"
	promotionStorage "api-billing/insfraestructure/postgres/promotion"

	"github.com/labstack/echo/v4"
)

const (
	privateRoutePrefix = "v1/invoice"
)

func NewRoutes(app *echo.Echo, db *sql.DB) {
	useCase := invoice.New(invoiceStorage.New(db))

	useCase.SetUseCasePromotion(promotion.New(promotionStorage.New(db)))
	useCase.SetUseCaseMedicine(medicine.New(medicineStorage.New(db)))

	handler := newHandler(useCase)

	privateRoutes(app, handler)
}

func privateRoutes(app *echo.Echo, handler handler) {
	api := app.Group(privateRoutePrefix)

	api.POST("", handler.create)
	api.GET("", handler.getAllByRangeCreatedAt)
	api.GET("/simulate", handler.GetPriceByMedicinesIDsAndSaleDate)
}
