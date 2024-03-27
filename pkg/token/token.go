package token

import (
	"fmt"
	"marketplace/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	AccesTokenExpiresAt   = time.Now().Add(15 * time.Minute).Unix()
	RefreshTokenExpiresAt = time.Now().Add(7 * 24 * time.Hour).Unix()
)

type Claims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateAccesToken(userID int, username string) (accessToken string, err error) {
	claims := &Claims{
		UserId:   fmt.Sprint(userID),
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: AccesTokenExpiresAt,
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(configs.JwtKey)
	if err != nil {
		return
	}

	return
}

func GenerateRefreshToken(userID int, username string) (refreshToken string, err error) {
	refreshClaims := &Claims{
		UserId:   fmt.Sprint(userID),
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: RefreshTokenExpiresAt,
		},
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(configs.RefreshJwtKey)
	if err != nil {
		return
	}

	return
}

func GenerateToken(userID int, username string) (accessToken string, refreshToken string, err error) {
	accessToken, err = GenerateAccesToken(userID, username)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = GenerateRefreshToken(userID, username)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
