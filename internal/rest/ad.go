package rest

import (
	"encoding/json"
	"marketplace/model"
	"marketplace/pkg/responseTemplate"
	"marketplace/pkg/serverErrors"
	"marketplace/pkg/utils"
	"marketplace/usecase"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
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
	contextLogger := utils.GetContextLogger(r.Context())

	decoder := json.NewDecoder(r.Body)
	adInput := new(model.AdAPI)
	err := decoder.Decode(adInput)
	if err != nil {
		contextLogger.WithFields(logrus.Fields{
			"error": err,
		}).
			Error("error while decoding request body")
		responseTemplate.ServeJsonError(w, serverErrors.ErrInvalidBody)
		return
	}

	ad, err := handler.adUsecase.AddAd(r.Context(), *adInput)
	if err != nil {
		contextLogger.WithFields(logrus.Fields{
			"error":   err,
			"adInput": adInput,
		}).
			Error("error while posting new ad")
		responseTemplate.ServeJsonError(w, err)
		return
	}

	responseTemplate.MarshalAndSend(w, ad)
}

func (handler *AdHandler) HandleGetAd(w http.ResponseWriter, r *http.Request) {
	contextLogger := utils.GetContextLogger(r.Context())

	pageNum, _ := strconv.Atoi(r.URL.Query().Get("page_num"))
	if pageNum < 1 {
		pageNum = 1
	}

	resultsPerPage, _ := strconv.Atoi(r.URL.Query().Get("results_per_page"))
	if resultsPerPage <= 0 {
		resultsPerPage = 10
	}

	sortField := r.URL.Query().Get("sort_by")
	if sortField != "date" && sortField != "price" {
		sortField = "created_at"
	} else if sortField == "price" {
		sortField = "cents_price"
	} else if sortField == "date" {
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
		contextLogger.WithFields(logrus.Fields{
			"error": err,
		}).
			Error("error while getting ads")
		responseTemplate.ServeJsonError(w, err)
		return
	}

	responseTemplate.MarshalAndSend(w, ads)
}
