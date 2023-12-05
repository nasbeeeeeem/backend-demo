package usecase

import (
	"backend-demo/pkg/domain/model"
	"backend-demo/pkg/domain/repository"
	"context"
	"time"
)

type UserUseCase interface {
	MeInfo(c context.Context, email string) (*model.User, error)
	Users(c context.Context) ([]*model.User, error)
	Update(c context.Context, id int, name string, email string, photoUrl string, accountCode string, bankCode string, branchCode string) (*model.User, error)
	Delete(c context.Context, id int) (*model.User, error)
}

type userUseCase struct {
	repository repository.UserRepository
	timeout    time.Duration
}

// MeInfo implements UserUseCase.
func (uu *userUseCase) MeInfo(c context.Context, email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()

	newUser, err := uu.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// Users implements UserUseCase.
func (uu *userUseCase) Users(c context.Context) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()

	user, err := uu.repository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update implements UserUseCase.
func (uu *userUseCase) Update(c context.Context, id int, name string, email string, photoUrl string, accountCode string, bankCode string, branchCode string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()

	updateUser, err := uu.repository.UpdateUser(ctx, id, name, email, photoUrl, accountCode, bankCode, branchCode)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

// Delete implements UserUseCase.
func (uu *userUseCase) Delete(c context.Context, id int) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()

	deleteUser, err := uu.repository.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return deleteUser, nil
}

func NewUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repository: userRepo,
	}
}
