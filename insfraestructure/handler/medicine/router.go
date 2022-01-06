package medicine

import (
	"database/sql"

	"api-billing/domain/medicine"
	medicineStorage "api-billing/insfraestructure/postgres/medicine"

	"github.com/labstack/echo/v4"
)

const (
	privateRoutePrefix = "v1/medicine"
)

func NewRoutes(app *echo.Echo, db *sql.DB) {
	useCase := medicine.New(medicineStorage.New(db))

	handler := newHandler(useCase)

	privateRoutes(app, handler)
}

func privateRoutes(app *echo.Echo, handler handler) {
	api := app.Group(privateRoutePrefix)

	api.POST("", handler.create)
	api.GET("", handler.getAll)
}
