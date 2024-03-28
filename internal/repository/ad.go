package repository

import (
	"context"
	"database/sql"
	"marketplace/model"
	"marketplace/pkg/queryTemplates"
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
		author_id,
		"name",
		description,
		cents_price,
		picture_url
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	var id int
	err := repo.db.QueryRow(
		query,
		ad.AuthorId,
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

func (repo *AdPg) GetAds(
	ctx context.Context,
	pageNum int,
	resultsPerPage int,
	sortField string,
	sortOrder string,
	minPrice string,
	maxPrice string,
) ([]model.AdDB, error) {
	query, args, err := queryTemplates.GenerateAdQuery(
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

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	adsToReturn := []model.AdDB{}

	for rows.Next() {
		var ad model.AdDB
		if err = rows.Scan(
			&ad.Id,
			&ad.AuthorId,
			&ad.Name,
			&ad.Description,
			&ad.CentsPrice,
			&ad.PictureUrl,
			&ad.CreatedAt,
		); err != nil {
			return nil, err
		}

		adsToReturn = append(adsToReturn, ad)
	}
	return adsToReturn, nil
}
