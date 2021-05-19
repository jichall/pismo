package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jichall/pismo/src/entities"
	"github.com/labstack/echo"
)

func createAccount(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	// parsing the body (as default bind is not available)
	c.Logger().Printf("%s", string(body))
	a := entities.Account{}

	err = json.Unmarshal(body, &a)

	if err != nil {
		c.Response().Writer.WriteHeader(http.StatusBadRequest)
		return err
	}

	c.Logger().Printf("%s", a.String())
	err = a.Insert(pool)

	if err != nil {
		c.Logger().Printf("%v", err)
		c.Response().Writer.WriteHeader(http.StatusInternalServerError)
	}

	return nil
}

func createTransaction(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	// parsing the body (as default bind is not available)
	c.Logger().Printf("%s", string(body))
	t := entities.Transaction{}

	err = json.Unmarshal(body, &t)

	c.Logger().Printf("%s", t.String())

	if err != nil {
		c.Logger().Printf("%v", err)
		return err
	}

	err = t.Insert(pool)

	if err != nil {
		c.Logger().Printf("%v", err)
	}

	return nil
}

func fetchAccount(c echo.Context) error {

	id := c.Param("account_id")
	parsed, err := strconv.ParseInt(id, 10, 0)

	if err != nil {
		return err
	}

	account := entities.Account{ID: int(parsed)}
	err = account.Select(pool)

	c.Logger().Printf(account.String())

	if err != nil {
		c.Logger().Printf("%v", err)
		return err
	}

	c.Logger().Printf("%v", account.String())
	c.JSON(http.StatusOK, account)

	return nil
}

func fetchTransaction(c echo.Context) error {
	return nil
}
