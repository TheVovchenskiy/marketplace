package serverErrors

import (
	"marketplace/internal/repository"
	"marketplace/pkg/listOptions"
	"marketplace/pkg/price"
	"marketplace/pkg/token"
	"marketplace/pkg/validator"
	"marketplace/usecase"
	"net/http"
)

var HTTPErrors = map[error]int{
	ErrMethodNotAllowed:                  http.StatusMethodNotAllowed,
	ErrInvalidRequest:                    http.StatusBadRequest,
	ErrInvalidBody:                       http.StatusBadRequest,
	validator.ErrInvalidPassword:         http.StatusBadRequest,
	validator.ErrInvalidUsername:         http.StatusBadRequest,
	token.ErrAuthorizationHeaderRequired: http.StatusUnauthorized,
	token.ErrInvalidToken:                http.StatusUnauthorized,
	repository.ErrAccountAlreadyExists:   http.StatusConflict,
	usecase.ErrInvalidLoginData:          http.StatusUnauthorized,
	price.ErrInvalidCents:                http.StatusBadRequest,
	price.ErrInvalidPriceFormat:          http.StatusBadRequest,
	listOptions.ErrInvalidQueryParam:     http.StatusBadRequest,
}

func MapHTTPError(err error) (msg string, status int) {
	if err == nil {
		err = ErrInternal
	}

	status, ok := HTTPErrors[err]
	if !ok {
		status = http.StatusInternalServerError
		err = ErrInternal
	}

	msg = err.Error()

	return
}
