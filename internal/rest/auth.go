package rest

import (
	"encoding/json"
	"marketplace/model"
	"marketplace/pkg/responseTemplate"
	"marketplace/pkg/serverErrors"
	"marketplace/pkg/utils"
	"marketplace/usecase"
	"net/http"

	"github.com/sirupsen/logrus"
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
	contextLogger := utils.GetContextLogger(r.Context())

	decoder := json.NewDecoder(r.Body)
	registrationInput := new(model.RegisterInput)
	err := decoder.Decode(registrationInput)
	if err != nil {
		contextLogger.WithFields(logrus.Fields{
			"error": err,
		}).
			Error("error while decoding request body")
		responseTemplate.ServeJsonError(w, serverErrors.ErrInvalidBody)
		return
	}

	user, err := handler.authUsecase.RegisterUser(*registrationInput)
	if err != nil {
		contextLogger.WithFields(logrus.Fields{
			"error": err,
		}).
			Error("error while registrating user")
		responseTemplate.ServeJsonError(w, err)
		return
	}

	responseTemplate.MarshalAndSend(w, user)
}

func (handler *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	contextLogger := utils.GetContextLogger(r.Context())

	decoder := json.NewDecoder(r.Body)
	loginInput := new(model.LoginInput)
	err := decoder.Decode(loginInput)
	if err != nil {
		contextLogger.WithFields(logrus.Fields{
			"error": err,
		}).
			Error("error while decoding request body")
		responseTemplate.ServeJsonError(w, serverErrors.ErrInvalidBody)
		return
	}

	user, err := handler.authUsecase.LoginUser(*loginInput)
	if err != nil {
		contextLogger.WithFields(logrus.Fields{
			"error": err,
		}).
			Error("error while logging user")
		responseTemplate.ServeJsonError(w, err)
		return
	}

	responseTemplate.MarshalAndSend(w, user)

}
