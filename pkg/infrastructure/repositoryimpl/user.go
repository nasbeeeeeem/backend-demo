package repositoryimpl

import (
	"backend-demo/pkg/domain/model"
	"backend-demo/pkg/domain/repository"
	"backend-demo/pkg/infrastructure/database"
	"context"
)

type userRepoImpl struct {
	DBClient *database.Engine
}

func NewUserRepoImpl(client *database.Engine) repository.UserRepository {
	return &userRepoImpl{
		DBClient: client,
	}
}

// 全ユーザーの取得
func (u *userRepoImpl) GetUsers(c context.Context) ([]*model.User, error) {
	var users []*model.User
	result := u.DBClient.Engine.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// Emailと一致するユーザーを取得
func (u *userRepoImpl) GetUserByEmail(c context.Context, email string) (*model.User, error) {
	var user *model.User
	resutl := u.DBClient.Engine.Where("email = ?", email).Find(&user)

	if resutl.Error != nil {
		return nil, resutl.Error
	}

	return user, nil
}

// プロフィール情報の更新
func (u *userRepoImpl) UpdateUser(c context.Context, id int, name string, email string, photoUrl string, accountCode string, bankCode string, branchCode string) (*model.User, error) {
	var updateUser *model.User
	// 更新対象のユーザー情報の取得
	// result := u.DBClient.Engine.First(&updateUser, id)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// 対象のユーザー情報の更新
	result := u.DBClient.Engine.Where("id = ?", id).Model(&updateUser).Updates(model.User{Name: name, Email: email, PhotoURL: photoUrl, AccountCode: accountCode, BankCode: bankCode, BranchCode: bankCode})
	if result.Error != nil {
		return nil, result.Error
	}

	return updateUser, nil
}

// ユーザーの削除
func (u *userRepoImpl) DeleteUser(c context.Context, id int) (*model.User, error) {
	var deleteUser *model.User

	result := u.DBClient.Engine.Delete(&deleteUser, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return deleteUser, nil
}
