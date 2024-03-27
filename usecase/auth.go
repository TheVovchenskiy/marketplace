package usecase

import (
	"context"
	"marketplace/model"
	"marketplace/pkg/token"
	"marketplace/pkg/validator"

	"github.com/google/uuid"
)

type AuthUsecase struct {
	userStorage UserStorage
}

func NewAuthUsecase(userStorage UserStorage) *AuthUsecase {
	return &AuthUsecase{
		userStorage: userStorage,
	}
}

func (u *AuthUsecase) RegisterUser(registerInput model.RegisterInput) (*model.User, error) {
	err := validator.ValidateRegisterInput(registerInput)
	if err != nil {
		return nil, err
	}

	salt := uuid.NewString()
	user := registerInput.ToUser(salt)

	user.Id, err = u.userStorage.StoreUser(context.Background(), user)
	if err != nil {
		return nil, err
	}

	user.AccessToken, err = token.GenerateAccesToken(user.Id, registerInput.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
