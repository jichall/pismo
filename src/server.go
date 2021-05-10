package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func serve(host, port string) {

	server := &http.Server{
		Addr:         host + ":" + port,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	e := echo.New()

	e.POST("/accounts", createAccount)
	e.POST("/transactions", createTransaction)

	e.GET("/accounts/:accountID", fetchAccount)

	e.Logger.Fatal(e.StartServer(server))
}
