package usecase

import (
	"backend-demo/ent"
	"backend-demo/pkg/domain/repository"
	"context"
	"time"
)

type UseCase interface {
	Create(c context.Context, name string, email string) (*ent.User, error)
	MeInfo(c context.Context, email string) (*ent.User, error)
	Users(c context.Context) ([]*ent.User, error)
}

type useCase struct {
	repository repository.UserRepository
	timeout    time.Duration
}

// Create implements UseCase.
func (uc *useCase) Create(c context.Context, name string, email string) (*ent.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()

	newUser, err := uc.repository.CreateUser(ctx, name, email)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// MeInfo implements UseCase.
func (uc *useCase) MeInfo(c context.Context, email string) (*ent.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()

	user, err := uc.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Users implements UseCase.
func (uc *useCase) Users(c context.Context) ([]*ent.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()

	users, err := uc.repository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func NewUseCase(userRepo repository.UserRepository) UseCase {
	return &useCase{
		repository: userRepo,
	}
}
