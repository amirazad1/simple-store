package service

import (
	"context"
	"database/sql"
	dbc "github.com/amirazad1/simple-store/db/sqlc"
)

type Store interface {
	dbc.Querier
	SaleTx(ctx context.Context, arg SaleTxParams) (int64, error)
}

type SQLStore struct {
	db *sql.DB
	*dbc.Queries
}

func NewStore(dbs *sql.DB) Store {
	return &SQLStore{
		db:      dbs,
		Queries: dbc.New(dbs),
	}
}
