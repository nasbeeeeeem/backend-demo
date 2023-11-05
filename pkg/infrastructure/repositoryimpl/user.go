package repositoryimpl

import (
	"backend-demo/pkg/domain/model"
	"backend-demo/pkg/domain/repository"
	"backend-demo/pkg/infrastructure/database"
	"context"
)

type userRepoImpl struct {
	Client *database.Client
}

func NewUserRepoImpl(client *database.Client) repository.UserRepository {
	return &userRepoImpl{
		Client: client,
	}
}

func (u *userRepoImpl) CreateUser(c context.Context, user *model.User) (*model.User, error) {
	panic("")
}

func (u *userRepoImpl) GetUsers(c context.Context) ([]*model.User, error) {
	panic("")
}

func (*userRepoImpl) GetUserByEmail(c context.Context, email string) (*model.User, error) {
	panic("")
}
