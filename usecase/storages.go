package usecase

import (
	"context"
	"marketplace/model"
)

type UserStorage interface {
	StoreUser(ctx context.Context, user *model.User) (int, error)
}

type SessionStorage interface {
}
