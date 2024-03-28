package repository

import (
	"context"
	"database/sql"
	"marketplace/model"
)

type AdPg struct {
	db *sql.DB
}

func NewAdPg(db *sql.DB) *AdPg {
	return &AdPg{
		db: db,
	}
}

func (repo *AdPg) AddAd(ctx context.Context, ad model.AdDB) (int, error) {
	query := `INSERT INTO public.ad (
		"name",
		description,
		cents_price,
		picture_url
	)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	var id int
	err := repo.db.QueryRow(
		query,
		ad.Name,
		ad.Description,
		ad.CentsPrice,
		ad.PictureUrl,
	).
		Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
