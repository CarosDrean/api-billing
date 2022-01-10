package invoice

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"api-billing/model"
)

type Invoice struct {
	storage Storage

	medicine  UseCaseMedicine
	promotion UseCasePromotion
}

func New(storage Storage) *Invoice {
	return &Invoice{storage: storage}
}

func (i *Invoice) SetUseCasePromotion(useCase UseCasePromotion) {
	i.promotion = useCase
}

func (i *Invoice) hasUseCasePromotion() bool {
	return i.promotion != nil
}

func (i *Invoice) SetUseCaseMedicine(useCase UseCaseMedicine) {
	i.medicine = useCase
}

func (i *Invoice) hasUseCaseMedicine() bool {
	return i.medicine != nil
}

func (i *Invoice) Create(m *model.Invoice) error {
	if err := i.validateDependencies(); err != nil {
		return fmt.Errorf("invoice.validateDependencies(): %v", err)
	}

	medicines, err := i.medicine.GetByIDs(m.MedicinesIDs)
	if err != nil {
		return fmt.Errorf("invoice.medicine.GetByIDs(): %v", err)
	}

	promotions, err := i.promotion.GetByStartAndFinishDate(time.Now(), time.Now())
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("invoice.promotion.GetByCreatedAt(): %v", err)
	}

	totalPrice := medicines.GetTotalPrice()

	m.TotalPrice = totalPrice - (totalPrice * (float64(promotions.GetBestPercentaje()) / 100))

	if err := i.storage.Create(m); err != nil {
		return fmt.Errorf("invoice.storage.create(): %v", err)
	}

	return nil
}

func (i *Invoice) GetAll() (model.Invoices, error) {
	if err := i.validateDependencies(); err != nil {
		return nil, fmt.Errorf("invoice.validateDependencies(): %v", err)
	}

	invoices, err := i.storage.GetAllWhere(model.Fields{}, model.SortFields{}, model.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("invoice.storage.GetAllWhere(): %v", err)
	}

	if err := i.buildDetails(invoices); err != nil {
		return nil, err
	}

	return invoices, nil
}

func (i *Invoice) GetAllByRangeCreatedAt(startDate, finishDate time.Time) (model.Invoices, error) {
	if err := i.validateDependencies(); err != nil {
		return nil, fmt.Errorf("invoice.validateDependencies(): %v", err)
	}

	invoices, err := i.storage.GetAllWhere(model.Fields{
		model.Field{Name: "created_at", Value: startDate, Operator: model.GreaterThanOrEqualTo},
		model.Field{Name: "created_at", Value: finishDate, Operator: model.LessThanOrEqualTo},
	}, model.SortFields{}, model.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("invoice.storage.GetAllWhere(): %v", err)
	}

	if err := i.buildDetails(invoices); err != nil {
		return nil, err
	}

	return invoices, nil
}

func (i *Invoice) GetPriceByMedicinesIDsAndSaleDate(medicinesIDs []uint, saleDate time.Time) (float64, error) {
	if err := i.validateDependencies(); err != nil {
		return 0, fmt.Errorf("invoice.validateDependencies(): %v", err)
	}

	medicines, err := i.medicine.GetByIDs(medicinesIDs)
	if err != nil {
		return 0, fmt.Errorf("invoice.medicine.GetByIDs(): %v", err)
	}

	promotions, err := i.promotion.GetByStartAndFinishDate(saleDate, saleDate)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("invoice.promotion.GetByCreatedAt(): %v", err)
	}

	totalPrice := medicines.GetTotalPrice()

	return totalPrice - (totalPrice * (float64(promotions.GetBestPercentaje()) / 100)), nil
}

func (i *Invoice) buildDetails(invoices model.Invoices) error {
	promotions, err := i.promotion.GetByIDs(invoices.GetPromotionIDs())
	if err != nil {
		return fmt.Errorf("invoice.promotion.GetByIDs(): %v", err)
	}

	medicines, err := i.medicine.GetByIDs(invoices.GetUniqueMedicinesIDs())
	if err != nil {
		return fmt.Errorf("invoice.medicine.GetByIDs(): %v", err)
	}

	for j, invoice := range invoices {
		invoices[j].Promotion, _ = promotions.FilterByID(invoice.PromotionID)

		for _, medicineID := range invoice.MedicinesIDs {
			medicine, _ := medicines.FilterByID(medicineID)
			invoices[j].Medicines = append(invoices[j].Medicines, medicine)
		}
	}

	return nil
}

func (i Invoice) validateDependencies() error {
	if !i.hasUseCaseMedicine() {
		return fmt.Errorf("%s", "dependency medicine not found")
	}

	if !i.hasUseCasePromotion() {
		return fmt.Errorf("%s", "dependency promotion not found")
	}

	return nil
}
