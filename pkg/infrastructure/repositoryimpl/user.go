package repositoryimpl

import (
	"backend-demo/ent"
	"backend-demo/ent/user"
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

// ユーザーの登録
func (u *userRepoImpl) CreateUser(c context.Context, name string, email string) (*ent.User, error) {
	newUser, err := u.Client.Client.User.Create().SetName(name).SetEmail(email).Save(context.Background())
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// 全ユーザーの取得
func (u *userRepoImpl) GetUsers(c context.Context) ([]*ent.User, error) {
	users, err := u.Client.Client.User.Query().Where().All(context.Background())
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Emailと一致するユーザーを取得
func (u *userRepoImpl) GetUserByEmail(c context.Context, email string) (*ent.User, error) {
	user, err := u.Client.Client.User.Query().Where(user.Email(email)).Only(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}
