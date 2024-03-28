package model

import (
	"marketplace/pkg/price"
	"strings"
)

type AdDB struct {
	Id          int
	AuthorId    int
	Name        string
	Description string
	CentsPrice  int64
	PictureUrl  string
	CreatedAt   string
}

func (ad *AdDB) ToAPI() (*AdAPI, error) {
	price, err := price.FromCents(ad.CentsPrice)
	if err != nil {
		return nil, err
	}
	return &AdAPI{
		Id:          ad.Id,
		AuthorId:    ad.AuthorId,
		Name:        ad.Name,
		Description: ad.Description,
		Price:       price,
		PictureUrl:  ad.PictureUrl,
		CreatedAt:   ad.CreatedAt,
	}, nil
}

type AdAPI struct {
	Id          int    `json:"id,omitempty"`
	AuthorId    int    `json:"authorId,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	PictureUrl  string `json:"pictureUrl"`
	CreatedAt   string `json:"createdAt"`
	MyAd        bool   `json:"myAd,omitempty"`
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
		AuthorId:    ad.AuthorId,
		Name:        ad.Name,
		Description: ad.Description,
		CentsPrice:  cents,
		PictureUrl:  ad.PictureUrl,
		CreatedAt:   ad.CreatedAt,
	}, nil
}
