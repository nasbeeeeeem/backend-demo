package database

import (
	"backend-demo/ent"
	"context"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

type Client struct {
	Client *ent.Client
}

func NewClient(dsn string) (*Client, error) {
	d := "postgresql://nasbeeeeeem:j1qzFMinVZY2@ep-yellow-snow-32012490.ap-southeast-1.aws.neon.tech/demo?sslmode=require"
	db, err := ent.Open(dialect.Postgres, d)
	if err != nil {
		return nil, err
	}
	//マイグレーション
	if err := db.Schema.Create(context.Background()); err != nil {
		db.Close()
		return nil, err
	}

	return &Client{Client: db}, nil
}
