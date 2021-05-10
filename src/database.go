package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseSettings struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func connect(settings DatabaseSettings) (*pgxpool.Pool, error) {

	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", settings.User,
		settings.Password, settings.Host, settings.Port, settings.Database)

	return pgxpool.Connect(context.Background(), url)
}
