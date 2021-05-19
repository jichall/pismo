package database

import (
	"log"

	"github.com/jackc/pgx/v4"
)

type PoolMock struct{}

func (p PoolMock) Exec(query string, args ...interface{}) error {
	log.Printf("PoolMock.Exec: %s", query)
	return nil
}

func (p PoolMock) Query(query string, args ...interface{}) (pgx.Rows, error) {
	log.Printf("PoolMock.Query: %s", query)
	return nil, nil
}
