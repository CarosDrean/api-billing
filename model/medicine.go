package model

import (
	"time"
)

type Medicine struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Medicines []Medicine

func (ms Medicines) FilterByID(ID uint) (Medicine, bool) {
	for _, medicine := range ms {
		if medicine.ID == ID {
			return medicine, true
		}
	}

	return Medicine{}, false
}

func (ms Medicines) GetTotalPrice() float64 {
	result := 0.0
	for _, medicine := range ms {
		result += medicine.Price
	}

	return result
}
