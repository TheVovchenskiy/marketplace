package routers

import (
	"marketplace/internal/rest"
	"marketplace/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func MountAuthRouter(router *mux.Router, userStorage usecase.UserStorage) {
	handler := rest.NewAuthHandler(userStorage)

	router.Handle("/register", http.HandlerFunc(handler.HandleRegistration)).Methods("POST")
	router.Handle("/login", http.HandlerFunc(handler.HandleLogin)).Methods("POST")
}
