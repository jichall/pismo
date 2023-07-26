package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jichall/pismo/src/entities"
	"github.com/labstack/echo"
)

func (api *API) CreateAccount(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	account := entities.Account{}

	err = json.Unmarshal(body, &account)

	if err != nil  || account.Validate() != nil {
		c.Response().Writer.WriteHeader(http.StatusBadRequest)
		return err
	}

	err = account.Insert(api.pool)

	if err != nil {
		c.Logger().Printf("%v", err)
		c.Response().Writer.WriteHeader(http.StatusInternalServerError)
	}

	return nil
}

func (api *API) GetAccount(c echo.Context) error {
	id := c.Param("account_id")
	parsed, err := strconv.ParseInt(id, 10, 0)

	if err != nil {
		return err
	}

	account := entities.Account{ID: int(parsed)}
	err = account.Select(api.pool)

	c.Logger().Printf(account.String())

	if err != nil {
		c.Logger().Printf("%v", err)
		return err
	}

	c.Logger().Printf("%v", account.String())
	c.JSON(http.StatusOK, account)

	return nil
}
