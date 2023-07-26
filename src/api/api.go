package api

import "github.com/jichall/pismo/src/database"

type API struct {
	pool *database.Pool
}

func New(pool *database.Pool) *API {
	return &API{
		pool: pool,
	}
}
