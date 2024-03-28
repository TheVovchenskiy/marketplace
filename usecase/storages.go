package usecase

import (
	"context"
	"marketplace/model"
)

type UserStorage interface {
	StoreUser(ctx context.Context, user *model.User) (int, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}

type AdStorage interface {
	AddAd(ctx context.Context, ad model.AdDB) (int, error)
	GetAds(
		ctx context.Context,
		pageNum int,
		resultsPerPage int,
		sortField string,
		sortOrder string,
		minPrice string,
		maxPrice string,
	) ([]model.AdDB, error)
}
