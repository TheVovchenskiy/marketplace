package routers

import (
	"marketplace/internal/rest"
	"marketplace/internal/rest/middleware"
	"marketplace/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func MountAdRouter(router *mux.Router, adStorage usecase.AdStorage) {
	handler := rest.NewAdHandler(adStorage)

	router.Handle("/ad", middleware.TokenVerify(http.HandlerFunc(handler.HandleAddAd))).Methods("POST")
}
