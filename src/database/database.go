package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	DatabaseSettings struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}

	// DatabaseFunctions is an interface that all instances of some database
	// must implement.
	DatabaseFunctions interface {
		Exec(query string, args ...interface{}) error
		Query(query string, args ...interface{}) (pgx.Rows, error)
	}
)

func Connect(settings DatabaseSettings) (*Pool, error) {

	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", settings.User,
		settings.Password, settings.Host, settings.Port, settings.Database)

	c, err := pgxpool.Connect(context.Background(), url)

	if err != nil {
		return nil, err
	}

	return &Pool{conn: c}, nil
}