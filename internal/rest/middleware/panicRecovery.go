package middleware

import (
	"marketplace/pkg/responseTemplate"
	"marketplace/pkg/serverErrors"
	"net/http"
)

func PanicRecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				responseTemplate.ServeJsonError(w, serverErrors.ErrInternal)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
