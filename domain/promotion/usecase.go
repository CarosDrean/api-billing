package promotion

import (
	"fmt"
	"time"

	"api-billing/model"
)

type Promotion struct {
	storage Storage
}

func New(storage Storage) *Promotion {
	return &Promotion{storage: storage}
}

func (p Promotion) Create(m *model.Promotion) error {
	promotion, err := p.storage.GetWhere(model.Fields{
		model.Field{Name: "start_date", Value: m.StartDate, Operator: model.GreaterThanOrEqualTo},
		model.Field{Name: "finish_date", Value: m.FinishDate, Operator: model.LessThanOrEqualTo},
	}, model.SortFields{})
	if err != nil {
		return err
	}

	if promotion.HasID() {
		return fmt.Errorf("promotion: the promotion will not be created because a promotion already exists in the range %v - %v", m.StartDate, m.FinishDate)
	}

	if err := p.storage.Create(m); err != nil {
		return fmt.Errorf("promotion.storage.create(): %v", err)
	}

	return nil
}

func (p Promotion) GetAll() (model.Promotions, error) {
	promotions, err := p.storage.GetAllWhere(model.Fields{}, model.SortFields{}, model.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("promotion.storage.GetAllWhere(): %v", err)
	}

	return promotions, nil
}

func (p Promotion) GetByIDs(IDs []uint) (model.Promotions, error) {
	promotions, err := p.storage.GetAllWhere(model.Fields{
		model.Field{Name: "id", Value: IDs, Operator: model.In},
	}, model.SortFields{}, model.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("promotion.storage.GetAllWhere(): %v", err)
	}

	return promotions, nil
}

func (p Promotion) GetByStartAndFinishDate(startDate, finishDate time.Time) (model.Promotions, error) {
	promotions, err := p.storage.GetAllWhere(model.Fields{
		model.Field{Name: "start_date", Value: startDate, Operator: model.GreaterThanOrEqualTo},
		model.Field{Name: "finish_date", Value: finishDate, Operator: model.LessThanOrEqualTo},
	}, model.SortFields{}, model.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("promotion.storage.GetAllWhere(): %v", err)
	}

	return promotions, nil
}
