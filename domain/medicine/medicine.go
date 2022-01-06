package medicine

import "api-billing/model"

type UseCase interface {
	Create(m *model.Medicine) error
	GetAll() (model.Medicines, error)
}

type Storage interface {
	Create(m *model.Medicine) error
	GetWhere(fields model.Fields, sortFields model.SortFields) (model.Medicine, error)
	GetAllWhere(fields model.Fields, sortFields model.SortFields, pagination model.Pagination) (model.Medicines, error)
}
