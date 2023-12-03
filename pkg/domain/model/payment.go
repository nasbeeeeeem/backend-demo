package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	EventID uint
	PaidBy  uint
	PaidTo  uint
	Amount  uint
}
