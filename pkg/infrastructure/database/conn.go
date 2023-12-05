package database

import (
	"backend-demo/pkg/domain/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Engine struct {
	Engine *gorm.DB
}

// DBのコネクション
func Conn(dsn string) (*Engine, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// dbのマイグレーション
	// db.AutoMigrate(&model.Bank{})
	db.AutoMigrate(&model.Bank{}, &model.User{}, &model.Group{}, &model.GroupUser{}, &model.Event{}, &model.Payment{})

	// サンプルデータの登録
	bank := model.Bank{
		ID:   "0005",
		Name: "三菱UFJ銀行",
	}

	db.Create(&bank)

	user1 := model.User{
		Name:  "nas",
		Email: "vividnasubi@gmail.com",
		// BankCode: bank.ID,
	}
	db.Create(&user1)

	group := model.Group{
		Name: "SampleGroup",
	}
	db.Create(&group)

	groupUser := model.GroupUser{
		UserID:  user1.ID,
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
		PaidBy:  user1.ID,
		PaidTo:  user1.ID,
		Amount:  1000,
	}
	db.Create(&payment)

	user2 := model.User{
		Name:  "rom",
		Email: "3103rom@gmail.com",
		// BankCode: bank.ID,
	}
	db.Create(&user2)

	groupUser2 := model.GroupUser{
		UserID:  user2.ID,
		GroupID: group.ID,
	}
	db.Create(&groupUser2)

	event2 := model.Event{
		Name:      "SampleEvent2",
		CreatedBy: groupUser2.UserID,
		GroupID:   groupUser2.GroupID,
	}
	db.Create(&event2)

	payment2 := model.Payment{
		EventID: event.ID,
		PaidBy:  user1.ID,
		PaidTo:  user1.ID,
		Amount:  2000,
	}
	db.Create(&payment2)

	fmt.Print("Sample data created successfully")

	return &Engine{Engine: db}, nil
}
