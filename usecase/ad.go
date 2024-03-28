package usecase

import (
	"context"
	"marketplace/model"
	"marketplace/pkg/validator"
)

type AdUsecase struct {
	adStorage AdStorage
}

func NewAdUsecase(adStorage AdStorage) *AdUsecase {
	return &AdUsecase{
		adStorage: adStorage,
	}
}

func (u *AdUsecase) AddAd(ad model.AdAPI) (model.AdAPI, error) {
	ad.Trim()

	if err := validator.ValidateAdAPI(ad); err != nil {
		return model.AdAPI{}, err
	}

	adDB, err := ad.ToDB()
	if err != nil {
		return model.AdAPI{}, err
	}

	ad.Id, err = u.adStorage.AddAd(context.TODO(), *adDB)
	if err != nil {
		return model.AdAPI{}, err
	}

	return ad, nil
}
