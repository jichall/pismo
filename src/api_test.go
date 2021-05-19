package main

import (
	"testing"

	"github.com/jichall/pismo/src/mock"
)

func TestCreateAccount(t *testing.T) {

	context := &mock.MockContext{}

	createAccount(context)
}
