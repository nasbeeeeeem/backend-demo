package database

import (
	"backend-demo/ent"
	// "backend-demo/pkg/infrastructure/cloudsql"
	"context"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

type Client struct {
	Client *ent.Client
}

func NewClient(dsn string) (*Client, error) {
	// dbPool, _ := cloudsql.ConnectWithConnector()
	db, err := ent.Open(dialect.Postgres, dsn)
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
