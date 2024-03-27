package usecase

import (
	"context"
	"marketplace/model"
)

type UserStorage interface {
	StoreUser(ctx context.Context, user *model.User) (int, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}
