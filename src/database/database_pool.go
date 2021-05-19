package database

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Pool is a type that encapsulates a connection pool
type Pool struct {
	conn *pgxpool.Pool
}

func (d *Pool) Exec(query string, args ...interface{}) error {

	conn, err := d.conn.Acquire(context.Background())

	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), query, args...)

	return err
}

func (d *Pool) Query(query string, args ...interface{}) (pgx.Rows, error) {

	conn, err := d.conn.Acquire(context.Background())

	if err != nil {
		return nil, err
	}
	defer conn.Release()

	r, err := conn.Query(context.Background(), query, args...)

	return r, err
}
