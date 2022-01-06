package medicine

import (
	"fmt"

	"api-billing/model"
)

type Medicine struct {
	storage Storage
}

func New(storage Storage) *Medicine {
	return &Medicine{storage: storage}
}

func (p Medicine) Create(m *model.Medicine) error {
	if err := p.storage.Create(m); err != nil {
		return fmt.Errorf("medicine.storage.create(): %v", err)
	}

	return nil
}

func (p Medicine) GetAll() (model.Medicines, error) {
	medicines, err := p.storage.GetAllWhere(model.Fields{}, model.SortFields{}, model.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("medicine.storage.GetAllWhere(): %v", err)
	}

	return medicines, nil
}

func (p Medicine) GetByIDs(IDs []uint) (model.Medicines, error) {
	medicines, err := p.storage.GetAllWhere(model.Fields{
		model.Field{Name: "id", Value: IDs, Operator: model.In},
	}, model.SortFields{}, model.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("medicine.storage.GetAllWhere(): %v", err)
	}

	return medicines, nil
}
