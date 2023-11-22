package repository

import (
	"backend-demo/ent"
	"context"
)

type UserRepository interface {
	CreateUser(c context.Context, name string, email string) (*ent.User, error)
	GetUsers(c context.Context) ([]*ent.User, error)
	GetUserByEmail(c context.Context, email string) (*ent.User, error)
	UpdateUser(c context.Context, name string, email string, photoUrl string, accountCode string, bankCode string, branchCode string) error
}
