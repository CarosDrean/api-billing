package promotion

import (
	"api-billing/domain/promotion"
	"api-billing/insfraestructure/handler/response"
	"api-billing/model"

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
