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

func NewClient() (*Client, error) {
	dsn := ""
	db, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Schema.Create(context.Background()); err != nil {
		db.Close()
		return nil, err
	}

	return &Client{Client: db}, nil
}
