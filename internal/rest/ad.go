package rest

import (
	"encoding/json"
	"fmt"
	"marketplace/model"
	"marketplace/pkg/responseTemplate"
	"marketplace/pkg/serverErrors"
	"marketplace/usecase"
	"net/http"
)

type AdHandler struct {
	authUsecase *usecase.AdUsecase
}

func NewAdHandler(adStorage usecase.AdStorage) *AdHandler {
	return &AdHandler{
		authUsecase: usecase.NewAdUsecase(adStorage),
	}
}

func (handler *AdHandler) HandleAddAd(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	adInput := new(model.AdAPI)
	err := decoder.Decode(adInput)
	if err != nil {
		responseTemplate.ServeJsonError(w, serverErrors.ErrInvalidBody)
		return
	}

	fmt.Println(adInput)

	ad, err := handler.authUsecase.AddAd(*adInput)
	if err != nil {
		responseTemplate.ServeJsonError(w, err)
		return
	}

	responseTemplate.MarshalAndSend(w, ad)
}
