package model

type User struct {
	Id           int    `json:"id,omitempty"`
	Username     string `json:"username,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	PasswordHash string `json:"-"`
	Salt         string `json:"-"`
}
