package postgres

import (
	"context"
	pgx "github.com/jackc/pgx/v5"
)

type DSN string

type Repo struct {
	db *pgx.Conn
}

func NewRepo(ctx context.Context, dsn DSN) (*Repo, error) {
	repo, err := pgx.Connect(ctx, string(dsn))
	return &Repo{db: repo}, err
}

func (r *Repo) Ping(ctx context.Context) error {
	return r.db.Ping(ctx)
}
