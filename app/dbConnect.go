package app

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"os"
)

func GetPostgres() (conn *sql.DB, err error) {
	for _, host := range []string{os.Getenv("PG_HOST"), "localhost"} {
		pgConnStr := fmt.Sprintf(
			"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
			os.Getenv("PG_USER"),
			os.Getenv("PG_DBNAME"),
			os.Getenv("PG_PASSWORD"),
			host,
			os.Getenv("PG_PORT"),
		)
		conn, err = sql.Open("pgx", pgConnStr)
		if err != nil {
			continue
		}
		err = conn.Ping()
		if err != nil {
			continue
		}

		break
	}

	return conn, err
}
