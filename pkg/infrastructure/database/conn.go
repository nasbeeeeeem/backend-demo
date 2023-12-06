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
	// localDsn := "host=localhost port=5432 user=postgres password=root dbname=gorm_db sslmode=disable"
	db, err := gorm.Open(postgres.Open("postgresql://nasbeeeeeem:j1qzFMinVZY2@ep-yellow-snow-32012490.ap-southeast-1.aws.neon.tech/demo?sslmode=require"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// dbの削除
	if err := db.Migrator().DropTable(&model.Bank{}, &model.User{}, &model.Group{}, &model.GroupUser{}, &model.Event{}, &model.Payment{}); err != nil {
		return nil, err
	}
	// dbのマイグレーション
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
	}
	db.Create(&user1)

	user2 := model.User{
		Name:  "rom",
		Email: "3103rom@gmail.com",
	}
	db.Create(&user2)

	group := model.Group{
		Name: "SampleGroup",
	}
	db.Create(&group)

	group2 := model.Group{
		Name: "SampleGroup2",
	}
	db.Create(&group2)

	groupUser1 := model.GroupUser{
		UserID:  user1.ID,
		GroupID: group.ID,
	}
	db.Create(&groupUser1)

	groupUser2 := model.GroupUser{
		UserID:  user2.ID,
		GroupID: group2.ID,
	}
	db.Create(&groupUser2)

	event1 := model.Event{
		Name:      "SampleEvent",
		CreatedBy: groupUser1.UserID,
		GroupID:   groupUser1.GroupID,
	}
	db.Create(&event1)

	event2 := model.Event{
		Name:      "SampleEvent2",
		CreatedBy: groupUser2.UserID,
		GroupID:   groupUser2.GroupID,
	}
	db.Create(&event2)

	payment1 := model.Payment{
		EventID: event1.ID,
		PaidBy:  user1.ID,
		PaidTo:  user2.ID,
		Amount:  1000,
	}
	db.Create(&payment1)

	payment2 := model.Payment{
		EventID: event2.ID,
		PaidBy:  user2.ID,
		PaidTo:  user1.ID,
		Amount:  2000,
	}
	db.Create(&payment2)

	fmt.Print("Sample data created successfully")

	return &Engine{Engine: db}, nil
}
