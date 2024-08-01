package postgres

import "github.com/jackc/pgx/v5"

type Repo struct {
	DB *pgx.Conn
}

func NewRepo(
	db *pgx.Conn,
) *Repo {
	return &Repo{
		DB: db,
	}
}
