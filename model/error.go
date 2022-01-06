package model

import "errors"

// Errors SQL
var (
	ErrUnique     = errors.New("Unique violation")
	ErrForeignKey = errors.New("Foreign key violation")
	ErrNotNull    = errors.New("Not Null violation")
)
