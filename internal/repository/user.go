package repository

import (
	"context"
	"database/sql"
	"marketplace/model"

	"github.com/jackc/pgx"
)

type UserPg struct {
	db *sql.DB
}

func NewUserPg(db *sql.DB) *UserPg {
	return &UserPg{
		db: db,
	}
}

func (repo *UserPg) StoreUser(ctx context.Context, user *model.User) (int, error) {
	query := `INSERT INTO public.user_profile (
		username,
		password_hash,
		salt
	)
	VALUES ($1, $2, $3)
	RETURNING id`

	var id int
	err := repo.db.QueryRow(
		query,
		user.Username,
		user.PasswordHash,
		user.Salt,
	).
		Scan(&id)

	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok {
			switch pgErr.Code {
			case "23505":
				return 0, ErrAccountAlreadyExists
			}
		}
		return 0, err
	}

	return id, nil
}
