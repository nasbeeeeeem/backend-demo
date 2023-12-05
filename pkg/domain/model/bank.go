package model

type Bank struct {
	ID    string `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Users []User `gorm:"foreignKey:BankCode;references:ID"`
}
