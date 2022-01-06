package promotion

import (
	"database/sql"

	"api-billing/domain/promotion"
	promotionStorage "api-billing/insfraestructure/postgres/promotion"

	"github.com/labstack/echo/v4"
)

const (
	privateRoutePrefix = "v1/promotion"
)

func NewRoutes(app *echo.Echo, db *sql.DB) {
	useCase := promotion.New(promotionStorage.New(db))

	handler := newHandler(useCase)

	privateRoutes(app, handler)
}

func privateRoutes(app *echo.Echo, handler handler) {
	api := app.Group(privateRoutePrefix)

	api.POST("", handler.create)
	api.GET("", handler.getAll)
}
