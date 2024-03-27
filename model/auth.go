package model

import "marketplace/pkg/hash"

type RegisterInput struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
}

func (i *RegisterInput) ToUser(salt string) *User {
	user := User{
		Username:     i.Username,
		PasswordHash: hash.HashPassword(i.Password, salt),
		Salt:         salt,
	}

	return &user
}
