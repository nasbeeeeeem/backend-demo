package database

import (
	"backend-demo/ent"
	"backend-demo/pkg/infrastructure/cloudsql"
	"context"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

type Client struct {
	Client *ent.Client
}

func NewClient(dsn string) (*Client, error) {
	dbPool, _ := cloudsql.ConnectWithConnector()
	drv := entsql.OpenDB(dialect.Postgres, dbPool)
	// defer drv.Close()

	// クライアントの初期化
	opt := []ent.Option{ent.Driver(drv)}
	db := ent.NewClient(opt...)

	//マイグレーション
	if err := db.Schema.Create(context.Background()); err != nil {
		// db.Close()
		return nil, err
	}

	return &Client{Client: db}, nil
}
