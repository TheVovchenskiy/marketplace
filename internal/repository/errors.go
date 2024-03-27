package repository

import "errors"

var (
	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrNoUserFound          = errors.New("no users found with given username")
)
