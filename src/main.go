package main

import (
	"flag"
	"log"

	"github.com/jichall/pismo/src/database"
)

var (
	host = flag.String("host", "127.0.0.1", "")
	port = flag.String("port", "8090", "")

	databaseSettings = database.DatabaseSettings{
		Host:     "localhost",
		Port:     "5432",
		User:     "pismo",
		Password: "a-very-strong-password",
		Database: "pismo",
	}

	pool *database.Pool
)

func main() {
	flag.Parse()

	// prevents variable shadowing
	var err error
	pool, err = database.Connect(databaseSettings)

	if err != nil {
		log.Fatalf("couldn't initialize connection pool, reason %v", err)
	}

	serve(*host, *port)
}
