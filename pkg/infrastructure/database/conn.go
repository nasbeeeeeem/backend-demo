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
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=root dbname=gorm_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"), &gorm.Config{})
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

	user := model.User{
		Name:  "nas",
		Email: "vividnasubi@gmail.com",
		// BankCode: bank.ID,
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

	return &Engine{Engine: db}, nil
}
