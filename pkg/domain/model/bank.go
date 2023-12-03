package model

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Name string `gorm:"not null"`
}
