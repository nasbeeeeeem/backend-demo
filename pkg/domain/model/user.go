package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string      `gorm:"not null"`
	Email          string      `gorm:"not null;unique"`
	GroupUsers     []GroupUser `gorm:"foreignKey:UserID"`
	PaymentsPaidBy []Payment   `gorm:"foreignKey:PaidBy"`
	PaymentsPaidTo []Payment   `gorm:"foreignKey:PaidTo"`
}
