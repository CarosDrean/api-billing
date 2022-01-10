package promotion

import (
	"api-billing/domain/promotion"
	"api-billing/insfraestructure/handler/response"
	"api-billing/model"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase promotion.UseCase
}

func newHandler(useCase promotion.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) create(c echo.Context) error {
	m := model.Promotion{}

	if err := c.Bind(&m); err != nil {
		return response.Failed("c.Bind()", response.BindFailed, err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return response.Failed("useCase.Create()", response.UnexpectedError, err)
	}

	return c.JSON(response.Successfull(response.RecordCreated, m))
}

func (h handler) getAll(c echo.Context) error {
	data, err := h.useCase.GetAll()
	if err != nil {
		return response.Failed("useCase.GetAll()", response.UnexpectedError, err)
	}

	return c.JSON(response.Successfull(response.Ok, data))
}

func (h handler) getByDate(c echo.Context) error {
	dateString := c.FormValue("date")

	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return response.Failed("time.Parse", response.Failure, err)
	}

	data, err := h.useCase.GetByDate(date)
	if errors.Is(err, sql.ErrNoRows) {
		return c.NoContent(http.StatusOK)
	}
	if err != nil {
		return response.Failed("useCase.GetByDate()", response.UnexpectedError, err)
	}

	return c.JSON(response.Successfull(response.Ok, data))
}
