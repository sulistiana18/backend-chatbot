package usecase

import (
	"context"
	"errors"

	"backend/entity"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo entity.UserRepository
}

func NewUserUsecase(repo entity.UserRepository) entity.UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userUsecase) CreateUser(ctx context.Context, name, email, password string) (*entity.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:       uuid.NewString(),
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}
	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Login(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}