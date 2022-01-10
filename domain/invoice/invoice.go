package invoice

import (
	"time"

	"api-billing/model"
)

type UseCase interface {
	Create(m *model.Invoice) error
	GetAll() (model.Invoices, error)
	GetAllByRangeCreatedAt(startDate, finishDate time.Time) (model.Invoices, error)
	GetPriceByMedicinesIDsAndSaleDate(medicinesIDs []uint, saleDate time.Time) (float64, error)
}

type Storage interface {
	Create(m *model.Invoice) error
	GetWhere(fields model.Fields, sortFields model.SortFields) (model.Invoice, error)
	GetAllWhere(fields model.Fields, sortFields model.SortFields, pagination model.Pagination) (model.Invoices, error)
}

type UseCaseMedicine interface {
	GetByIDs(IDs []uint) (model.Medicines, error)
}

type UseCasePromotion interface {
	GetByIDs(IDs []uint) (model.Promotions, error)
	GetByStartAndFinishDate(startDate, finishDate time.Time) (model.Promotions, error)
}
