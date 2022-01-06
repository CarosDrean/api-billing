package medicine

import (
	"fmt"
	"net/http"

	"api-billing/domain/medicine"
	"api-billing/model"

	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase medicine.UseCase
}

func newHandler(useCase medicine.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) create(c echo.Context) error {
	m := model.Medicine{}

	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("c.Bind()"))
	}

	if err := h.useCase.Create(&m); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err).Error())
	}

	return c.JSON(http.StatusCreated, fmt.Sprintf("created successful"))
}

func (h handler) getAll(c echo.Context) error {
	data, err := h.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err).Error())
	}

	return c.JSON(http.StatusOK, data)
}
