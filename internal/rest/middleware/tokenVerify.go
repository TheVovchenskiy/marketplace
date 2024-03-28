package middleware

import (
	"context"
	"fmt"
	"marketplace/configs"
	"marketplace/pkg/responseTemplate"
	"marketplace/pkg/token"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func TokenVerify(strict bool, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			if !strict {
				next.ServeHTTP(w, r)
				return
			}
			responseTemplate.ServeJsonError(w, token.ErrAuthorizationHeaderRequired)
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			responseTemplate.ServeJsonError(w, token.ErrInvalidToken)
			return
		}

		tokenString := bearerToken[1]
		claims := &token.Claims{}

		accessToken, err := jwt.ParseWithClaims(tokenString, claims, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("%w: %v", token.ErrUnexpectedSigningMethod, jwtToken.Header["alg"])
			}
			return configs.JwtKey, nil
		})

		if err != nil {
			responseTemplate.ServeJsonError(w, token.ErrInvalidToken)
			return
		}

		if !accessToken.Valid {
			responseTemplate.ServeJsonError(w, token.ErrInvalidToken)
			return
		}

		ctx := context.WithValue(r.Context(), token.UserContextKey, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
