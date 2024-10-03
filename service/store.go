package service

import (
	"database/sql"
	dbc "github.com/amirazad1/simple-store/db/sqlc"
)

type Store struct {
	db *sql.DB
	*dbc.Queries
}

func NewStore(dbs *sql.DB) *Store {
	return &Store{
		db:      dbs,
		Queries: dbc.New(dbs),
	}
}
