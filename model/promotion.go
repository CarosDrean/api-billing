package model

import (
	"time"
)

type Promotion struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	Percentage  int       `json:"percentage"`
	StartDate   time.Time `json:"start_date"`
	FinishDate  time.Time `json:"finish_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p Promotion) HasID() bool {
	return p.ID != 0
}

type Promotions []Promotion

func (ps Promotions) FilterByID(ID uint) (Promotion, bool) {
	for _, promotion := range ps {
		if promotion.ID == ID {
			return promotion, true
		}
	}

	return Promotion{}, false
}

func (ps Promotions) GetBestPercentaje() int {
	bestPromotion := 0
	for _, promotion := range ps {
		if promotion.Percentage > bestPromotion {
			bestPromotion = promotion.Percentage
		}
	}

	return bestPromotion
}
