package usecase

import (
	"context"
	"errors"
	"marketplace/internal/repository"
	"marketplace/model"
	"marketplace/pkg/hash"
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
	registerInput.Trim()
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

func (u *AuthUsecase) LoginUser(loginInput model.LoginInput) (*model.User, error) {
	loginInput.Trim()
	user, err := u.userStorage.GetUserByUsername(context.Background(), loginInput.Username)
	if err != nil {
		if errors.Is(err, repository.ErrNoUserFound) {
			err = ErrInvalidLoginData
		}
		return nil, err
	}

	if !hash.MatchPasswords(user.PasswordHash, loginInput.Password, user.Salt) {
		return nil, ErrInvalidLoginData
	}

	user.AccessToken, err = token.GenerateAccesToken(user.Id, loginInput.Username)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
