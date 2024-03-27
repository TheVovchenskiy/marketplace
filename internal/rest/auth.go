package rest

import (
	"encoding/json"
	"marketplace/model"
	"marketplace/pkg/responseTemplate"
	"marketplace/pkg/serverErrors"
	"marketplace/usecase"
	"net/http"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthHandler(authStorage usecase.UserStorage) *AuthHandler {
	return &AuthHandler{
		authUsecase: usecase.NewAuthUsecase(authStorage),
	}
}

func (handler *AuthHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	registrationInput := new(model.RegisterInput)
	err := decoder.Decode(registrationInput)
	if err != nil {
		responseTemplate.ServeJsonError(w, serverErrors.ErrInvalidBody)
		return
	}


	user, err := handler.authUsecase.RegisterUser(*registrationInput)
	if err != nil {
		responseTemplate.ServeJsonError(w, err)
		return
	}

	responseTemplate.MarshalAndSend(w, user)
}
