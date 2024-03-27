package model

import (
	"marketplace/pkg/hash"
	"strings"
)

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (i *RegisterInput) Trim() {
	i.Username = strings.TrimSpace(i.Username)
}

func (i *RegisterInput) ToUser(salt string) *User {
	user := User{
		Username:     i.Username,
		PasswordHash: hash.HashPassword(i.Password, salt),
		Salt:         salt,
	}

	return &user
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (i *LoginInput) Trim() {
	i.Username = strings.TrimSpace(i.Username)
}
