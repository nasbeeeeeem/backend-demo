package database

import (
	"backend-demo/pkg/domain/model"
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDB() {
	// dbのマイグレーション
	db.AutoMigrate(&model.User{}, &model.Group{}, &model.GroupUser{}, &model.Event{}, &model.Payment{})

	// サンプルデータの登録
	user := model.User{
		Name:  "yakiu",
		Email: "yakiu@gmail.com",
	}
	db.Create(&user)

	group := model.Group{
		Name: "SampleGroup",
	}
	db.Create(&group)

	groupUser := model.GroupUser{
		UserID:  user.ID,
		GroupID: group.ID,
	}
	db.Create(&groupUser)

	event := model.Event{
		Name:      "SampleEvent",
		CreatedBy: groupUser.UserID,
		GroupID:   groupUser.GroupID,
	}
	db.Create(&event)

	payment := model.Payment{
		EventID: event.ID,
		PaidBy:  user.ID,
		PaidTo:  user.ID,
		Amount:  1000,
	}
	db.Create(&payment)

	fmt.Print("Sample data created successfully")
}
