package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `gorm:"not null"`
	Email          string `gorm:"not null;unique"`
	PhotoURL       string
	AccountCode    string
	BankCode       string `gorm:"default:null"`
	BranchCode     string
	GroupUsers     []GroupUser `gorm:"foreignKey:UserID"`
	PaymentsPaidBy []Payment   `gorm:"foreignKey:PaidBy"`
	PaymentsPaidTo []Payment   `gorm:"foreignKey:PaidTo"`
}
