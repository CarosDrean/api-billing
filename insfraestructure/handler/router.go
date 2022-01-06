package handler

import (
	"database/sql"

	"api-billing/insfraestructure/handler/invoice"
	"api-billing/insfraestructure/handler/medicine"
	"api-billing/insfraestructure/handler/promotion"

	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Echo, db *sql.DB) {
	invoice.NewRoutes(app, db)
	medicine.NewRoutes(app, db)
	promotion.NewRoutes(app, db)
}
