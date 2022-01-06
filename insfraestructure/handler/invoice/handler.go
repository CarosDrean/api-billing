package invoice

import (
	"fmt"
	"net/http"
	"time"

	"api-billing/domain/invoice"
	"api-billing/model"

	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase invoice.UseCase
}

func newHandler(useCase invoice.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) create(c echo.Context) error {
	m := model.Invoice{}

	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("c.Bind()"))
	}

	if err := h.useCase.Create(&m); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err).Error())
	}

	return c.JSON(http.StatusCreated, fmt.Sprintf("created successful"))
}

func (h handler) getAllByRangeCreatedAt(c echo.Context) error {
	startDateString := c.FormValue("start-date")
	finishDateString := c.FormValue("finish-date")

	startDate, err := time.Parse(time.RFC3339, startDateString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("time.Parse()"))
	}

	finishDate, err := time.Parse(time.RFC3339, finishDateString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("time.Parse()"))
	}

	data, err := h.useCase.GetAllByRangeCreatedAt(startDate, finishDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err).Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h handler) GetPriceByMedicinesIDsAndSaleDate(c echo.Context) error {
	m := model.Invoice{}

	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("c.Bind()"))
	}

	saleDateString := c.FormValue("sale-date")
	saleDate, err := time.Parse(time.RFC3339, saleDateString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("time.Parse()"))
	}

	data, err := h.useCase.GetPriceByMedicinesIDsAndSaleDate(m.MedicinesIDs, saleDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err).Error())
	}

	return c.JSON(http.StatusOK, data)
}
