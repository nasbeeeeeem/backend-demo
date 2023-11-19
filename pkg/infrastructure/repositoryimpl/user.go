package repositoryimpl

import (
	"backend-demo/ent"
	"backend-demo/ent/user"
	"backend-demo/pkg/domain/repository"
	"backend-demo/pkg/infrastructure/database"
	"context"
	"fmt"
)

type userRepoImpl struct {
	DBClient *database.Client
}

func NewUserRepoImpl(client *database.Client) repository.UserRepository {
	return &userRepoImpl{
		DBClient: client,
	}
}

// トランザクション制御のヘルパー関数
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

// ユーザーの登録
func (u *userRepoImpl) CreateUser(c context.Context, name string, email string) (*ent.User, error) {
	// err := WithTx(c, u.DBClient.Client, func(tx *ent.Tx) error {
	// 	user, err := tx.User.Create().SetName(name).SetEmail(email).Save(c)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })

	// if err != nil {
	// 	return nil, err
	// }

	// return user, nil

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
