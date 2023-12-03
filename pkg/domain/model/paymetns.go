package model

import (
	"gorm.io/gorm"
)

type Payments struct {
	gorm.Model
	EventID uint
	PaidBy  uint
	PaidTo  uint
	GroupID uint
	Amount  uint
}
