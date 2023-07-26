package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jichall/pismo/src/entities"
	"github.com/labstack/echo"
)

func (api *API) CreateTransaction(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	transaction := entities.Transaction{}

	err = json.Unmarshal(body, &transaction)

	if err != nil {
		c.Logger().Printf("%v", err)
		return err
	}

	c.Logger().Printf("%s", transaction.String())

	err = transaction.Insert(api.pool)

	if err != nil {
		c.Logger().Printf("%v", err)
		c.Response().Writer.WriteHeader(http.StatusInternalServerError)
	}

	return nil
}

func (api *API) GetTransaction(c echo.Context) error {
	return nil
}
