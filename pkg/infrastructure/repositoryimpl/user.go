package repositoryimpl

import (
	"backend-demo/ent"
	"backend-demo/ent/user"
	"backend-demo/pkg/domain/repository"
	"backend-demo/pkg/infrastructure/database"
	"context"
)

type userRepoImpl struct {
	DBClient *database.Client
}

func NewUserRepoImpl(client *database.Client) repository.UserRepository {
	return &userRepoImpl{
		DBClient: client,
	}
}

// ユーザーの登録
func (u *userRepoImpl) CreateUser(c context.Context, name string, email string) (*ent.User, error) {
	tx, err := u.DBClient.Client.Tx(c)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	newUser, err := tx.User.Create().SetName(name).SetEmail(email).Save(c)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = rerr
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return newUser, nil
}

// 全ユーザーの取得
func (u *userRepoImpl) GetUsers(c context.Context) ([]*ent.User, error) {
	users, err := u.DBClient.Client.User.Query().Where().All(context.Background())
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Emailと一致するユーザーを取得
func (u *userRepoImpl) GetUserByEmail(c context.Context, email string) (*ent.User, error) {
	user, err := u.DBClient.Client.User.Query().Where(user.Email(email)).Only(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}
