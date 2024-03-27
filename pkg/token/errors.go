package token

import "errors"

var (
	ErrAuthorizationHeaderRequired = errors.New("authorization header is required")
	ErrInvalidToken                = errors.New("invalid authorization token")
	ErrUnexpectedSigningMethod     = errors.New("unexpected signing method")
)
