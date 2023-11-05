package repossitory

import (
	"backend-demo/pkg/domain/model"
	"context"
)

type UserRepository interface {
	CreateUser(c context.Context, user *model.User) (*model.User, error)
	GetUsers(c context.Context) ([]*model.User, error)
	GetUserByEmail(c context.Context, email string) (*model.User, error)
}
