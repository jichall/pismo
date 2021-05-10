package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func createAccount(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	// parsing the body (as default bind is not available)
	c.Logger().Printf("%s", string(body))
	a := Account{}

	err = json.Unmarshal(body, &a)

	if err != nil {
		c.Response().Writer.WriteHeader(http.StatusBadRequest)
		return err
	}

	err = a.Insert()

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
	t := Transaction{}

	err = json.Unmarshal(body, &t)

	if err != nil {
		c.Logger().Printf("%v", err)
		c.Response().Writer.WriteHeader(http.StatusBadRequest)
		return err
	}

	err = t.Insert()

	if err != nil {
		c.Logger().Printf("%v", err)
		c.Response().Writer.WriteHeader(http.StatusInternalServerError)
	}

	return nil
}

func fetchAccount(c echo.Context) error {

	id := c.Param("account_id")
	parsed, err := strconv.ParseInt(id, 10, 0)

	if err != nil {
		return err
	}

	account := Account{ID: int(parsed)}
	err = account.Select()

	if err != nil {
		c.Logger().Printf("%v", err)
		return err
	}

	c.Logger().Printf("%v", account.String())
	c.JSON(http.StatusOK, account)

	return nil
}
