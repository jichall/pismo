package main

import (
	"net/http"
	"time"

	"github.com/jichall/pismo/src/api"
	"github.com/labstack/echo"
)

func serve(host, port string) {

	server := &http.Server{
		Addr:         host + ":" + port,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	e := echo.New()

	a := api.New(pool)

	e.POST("/accounts", a.CreateAccount)
	e.POST("/transactions", a.CreateTransaction)

	e.GET("/accounts/:account_id", a.GetAccount)
	e.GET("/transactions/:transaction_id", a.GetTransaction)

	e.Logger.Fatal(e.StartServer(server))
}
