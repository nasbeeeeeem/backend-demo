package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string      `gorm:"not null"`
	Email          string      `gorm:"not null;unique"`
	GroupUsers     []GroupUser `gorm:"foreignKey:UserID"`
	PaymentsPaidBy []Payments  `gorm:"foreignKey:PaidBy"`
	PaymentsPaidTo []Payments  `gorm:"foreignKey:PaidTo"`
}
