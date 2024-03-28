package model

import (
	"marketplace/pkg/price"
	"strings"
)

type AdDB struct {
	Id          int
	Name        string
	Description string
	CentsPrice  int64
	PictureUrl  string
}

func (ad *AdDB) ToAPI() (*AdAPI, error) {
	price, err := price.FromCents(ad.CentsPrice)
	if err != nil {
		return nil, err
	}
	return &AdAPI{
		Id:          ad.Id,
		Name:        ad.Name,
		Description: ad.Description,
		Price:       price,
		PictureUrl:  ad.PictureUrl,
	}, nil
}

type AdAPI struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	PictureUrl  string `json:"pictureUrl"`
}

func (ad *AdAPI) Trim() {
	ad.Name = strings.TrimSpace(ad.Name)
	ad.Description = strings.TrimSpace(ad.Description)
	ad.Price = strings.TrimSpace(ad.Price)
	ad.PictureUrl = strings.TrimSpace(ad.PictureUrl)
}

func (ad *AdAPI) ToDB() (*AdDB, error) {
	cents, err := price.ToCents(ad.Price)
	if err != nil {
		return nil, err
	}
	return &AdDB{
		Id:          ad.Id,
		Name:        ad.Name,
		Description: ad.Description,
		CentsPrice:  cents,
		PictureUrl:  ad.PictureUrl,
	}, nil
}
