package repository

import (
	"context"
	"database/sql"
	"fmt"
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

func (repo *UserPg) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	query := `SELECT
					up.id,
					up.username,
					up.password_hash,
					up.salt
				FROM
					public.user_profile up
				WHERE
					up.username = $1
				`
	var user model.User
	err := repo.db.QueryRow(
		query,
		username,
	).
		Scan(
			&user.Id,
			&user.Username,
			&user.PasswordHash,
			&user.Salt,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("%w: %s", ErrNoUserFound, username)
		}
		return model.User{}, err
	}

	return user, nil
}
