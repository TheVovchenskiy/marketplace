package rest

import (
	"encoding/json"
	"fmt"
	"marketplace/model"
	"marketplace/pkg/responseTemplate"
	"marketplace/pkg/serverErrors"
	"marketplace/usecase"
	"net/http"
	"strconv"
)

type AdHandler struct {
	adUsecase *usecase.AdUsecase
}

func NewAdHandler(adStorage usecase.AdStorage) *AdHandler {
	return &AdHandler{
		adUsecase: usecase.NewAdUsecase(adStorage),
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

	ad, err := handler.adUsecase.AddAd(r.Context(), *adInput)
	if err != nil {
		responseTemplate.ServeJsonError(w, err)
		return
	}

	responseTemplate.MarshalAndSend(w, ad)
}

func (handler *AdHandler) HandleGetAd(w http.ResponseWriter, r *http.Request) {
	pageNum, _ := strconv.Atoi(r.URL.Query().Get("page_num"))
	if pageNum < 1 {
		pageNum = 1
	}

	resultsPerPage, _ := strconv.Atoi(r.URL.Query().Get("results_per_page"))
	if resultsPerPage <= 0 {
		resultsPerPage = 10
	}

	sortField := r.URL.Query().Get("sort_by")
	if sortField != "created_at" && sortField != "cents_price" {
		sortField = "created_at"
	}

	sortOrder := r.URL.Query().Get("order")
	if sortOrder != "desc" && sortOrder != "asc" {
		sortOrder = "desc"
	}

	minPrice := r.URL.Query().Get("min_price")

	maxPrice := r.URL.Query().Get("max_price")

	ads, err := handler.adUsecase.GetAds(
		r.Context(),
		pageNum,
		resultsPerPage,
		sortField,
		sortOrder,
		minPrice,
		maxPrice,
	)
	if err != nil {
		fmt.Println(err)
		responseTemplate.ServeJsonError(w, err)
		return
	}

	responseTemplate.MarshalAndSend(w, ads)
}
