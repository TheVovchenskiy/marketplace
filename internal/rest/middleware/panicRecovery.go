package middleware

import (
	"marketplace/pkg/responseTemplate"
	"marketplace/pkg/serverErrors"
	"net/http"

	"github.com/sirupsen/logrus"
)

func PanicRecoverMiddleware(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					responseTemplate.ServeJsonError(w, serverErrors.ErrInternal)

					logger.WithFields(logrus.Fields{
						"status":   http.StatusInternalServerError,
						"method":   r.Method,
						"URL":      r.URL.Path,
						"endpoint": r.RemoteAddr,
						"panic":    err,
					}).Error("recovered")
					responseTemplate.ServeJsonError(w, serverErrors.ErrInternal)

					return
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
