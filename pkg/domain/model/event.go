package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name            string `gorm:"not null"`
	CreatedBy       uint
	GroupID         uint
	PaymentsEventID []Payments `gorm:"foreignKey:EventID"`
}
