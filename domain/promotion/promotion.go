package promotion

import (
	"time"

	"api-billing/model"
)

type UseCase interface {
	Create(m *model.Promotion) error
	GetAll() (model.Promotions, error)
	GetByDate(date time.Time) (model.Promotion, error)
}

type Storage interface {
	Create(m *model.Promotion) error
	GetWhere(fields model.Fields, sortFields model.SortFields) (model.Promotion, error)
	GetAllWhere(fields model.Fields, sortFields model.SortFields, pagination model.Pagination) (model.Promotions, error)
}
