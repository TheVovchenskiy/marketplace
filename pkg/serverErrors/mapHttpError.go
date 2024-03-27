package serverErrors

import (
	"marketplace/pkg/token"
	"marketplace/pkg/validator"
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
