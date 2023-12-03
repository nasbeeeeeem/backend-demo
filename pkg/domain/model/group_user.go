package model

type GroupUser struct {
	UserID        uint    `gorm:"unique"`
	GroupID       uint    `gorm:"unique"`
	EventsUserID  []Event `gorm:"foreignKey:CreatedBy;references:UserID"` // CreatedBy列の外部キー
	EventsGroupID []Event `gorm:"foreignKey:GroupID;references:GroupID"`  // GroupID列の外部キー
}
