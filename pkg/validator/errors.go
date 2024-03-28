package validator

import "errors"

var (
	ErrInvalidUsername      = errors.New("invalid username")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrInvalidAdName        = errors.New("invalid ad name")
	ErrInvalidAdDescription = errors.New("invalid ad description")
)
