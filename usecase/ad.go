package usecase

import (
	"context"
	"errors"
	"fmt"
	"marketplace/model"
	"marketplace/pkg/token"
	"marketplace/pkg/utils"
	"marketplace/pkg/validator"
	"strconv"
)

type AdUsecase struct {
	adStorage AdStorage
}

func NewAdUsecase(adStorage AdStorage) *AdUsecase {
	return &AdUsecase{
		adStorage: adStorage,
	}
}

func (u *AdUsecase) AddAd(ctx context.Context, ad model.AdAPI) (model.AdAPI, error) {
	ad.Trim()

	authorId, err := utils.GetContextUserId(ctx)
	if err != nil {
		if errors.Is(err, utils.ErrNoUserIdInContext) {
			return model.AdAPI{}, token.ErrInvalidToken
		}
		return model.AdAPI{}, err
	}
	ad.AuthorId = authorId

	if err := validator.ValidateAdAPI(ad); err != nil {
		return model.AdAPI{}, err
	}

	adDB, err := ad.ToDB()
	if err != nil {
		return model.AdAPI{}, err
	}

	ad.Id, err = u.adStorage.AddAd(context.TODO(), *adDB)
	if err != nil {
		fmt.Println(err)
		return model.AdAPI{}, err
	}

	return ad, nil
}

func (u *AdUsecase) GetAds(
	ctx context.Context,
	pageNum int,
	resultsPerPage int,
	sortField string,
	sortOrder string,
	minPrice string,
	maxPrice string,
) ([]model.AdAPI, error) {
	var userIdStr string
	userId, err := utils.GetContextUserId(ctx)
	if err != nil {
		if errors.Is(err, utils.ErrNoUserIdInContext) {
			userIdStr = ""
		} else {
			return nil, err
		}
	} else {
		userIdStr = fmt.Sprint(userId)
	}

	ads, err := u.adStorage.GetAds(
		ctx,
		pageNum,
		resultsPerPage,
		sortField,
		sortOrder,
		minPrice,
		maxPrice,
	)
	if err != nil {
		return nil, err
	}

	adsToReturn := []model.AdAPI{}
	for _, ad := range ads {
		adApi, err := ad.ToAPI()
		if err != nil {
			return nil, err
		}

		if len(userIdStr) > 0 {
			userId, _ := strconv.Atoi(userIdStr)
			adApi.MyAd = userId == ad.AuthorId
		} else {
			adApi.MyAd = false
		}

		adsToReturn = append(adsToReturn, *adApi)
	}
	return adsToReturn, nil
}
