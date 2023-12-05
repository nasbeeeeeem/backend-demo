package repository

import (
	"backend-demo/pkg/domain/model"
	"context"
)

type UserRepository interface {
	GetUsers(c context.Context) ([]*model.User, error)
	GetUserByEmail(c context.Context, email string) (*model.User, error)
	UpdateUser(c context.Context, id int, name string, email string, photoUrl string, accountCode string, bankCode string, branchCode string) (*model.User, error)
	DeleteUser(c context.Context, id int) (*model.User, error)
}
