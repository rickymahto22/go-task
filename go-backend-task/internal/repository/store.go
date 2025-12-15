package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	db "go-backend-task/db/sqlc"
)

type Store struct {
	*db.Queries
	Pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		Queries: db.New(pool),
		Pool:    pool,
	}
}