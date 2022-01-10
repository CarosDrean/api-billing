package model

import (
	"encoding/json"
	"time"
)

type Invoice struct {
	ID           uint      `json:"id"`
	TotalPrice   float64   `json:"total_price"`
	PromotionID  uint      `json:"promotion_id"`
	Promotion    Promotion `json:"promotion"`
	MedicinesIDs []uint    `json:"medicines_ids"`
	Medicines    Medicines `json:"medicines"`
	CreatedAt    time.Time `json:"created_at"`

	UpdatedAt           time.Time        `json:"updated_at"`
	MedicinesIDsRawJSON *json.RawMessage `json:"-"`
}

type Invoices []Invoice

func (is Invoices) GetPromotionIDs() []uint {
	ids := make([]uint, 0)
	for _, invoice := range is {
		ids = append(ids, invoice.PromotionID)
	}

	return UniqueUints(ids)
}

func (is Invoices) GetUniqueMedicinesIDs() []uint {
	ids := make([]uint, 0)
	for _, invoice := range is {
		ids = append(ids, invoice.MedicinesIDs...)
	}

	return UniqueUints(ids)
}
