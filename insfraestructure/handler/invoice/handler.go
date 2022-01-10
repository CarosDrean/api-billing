package invoice

import (
	"time"

	"api-billing/domain/invoice"
	"api-billing/insfraestructure/handler/response"
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

func (h handler) getAllByRangeCreatedAt(c echo.Context) error {
	startDateString := c.FormValue("start-date")
	finishDateString := c.FormValue("finish-date")

	startDate, err := time.Parse(time.RFC3339, startDateString)
	if err != nil {
		return response.Failed("time.Parse", response.Failure, err)
	}

	finishDate, err := time.Parse(time.RFC3339, finishDateString)
	if err != nil {
		return response.Failed("time.Parse", response.Failure, err)
	}

	data, err := h.useCase.GetAllByRangeCreatedAt(startDate, finishDate)
	if err != nil {
		return response.Failed("useCase.GetAllByRangeCreatedAt()", response.UnexpectedError, err)
	}

	return c.JSON(response.Successfull(response.Ok, data))
}

func (h handler) GetPriceByMedicinesIDsAndSaleDate(c echo.Context) error {
	m := model.Invoice{}

	if err := c.Bind(&m); err != nil {
		return response.Failed("c.Bind()", response.BindFailed, err)
	}

	saleDateString := c.FormValue("sale-date")
	saleDate, err := time.Parse(time.RFC3339, saleDateString)
	if err != nil {
		return response.Failed("time.Parse", response.Failure, err)
	}

	data, err := h.useCase.GetPriceByMedicinesIDsAndSaleDate(m.MedicinesIDs, saleDate)
	if err != nil {
		return response.Failed("useCase.GetPriceByMedicinesIDsAndSaleDate()", response.UnexpectedError, err)
	}

	return c.JSON(response.Successfull(response.Ok, data))
}
