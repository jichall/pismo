package entities

import (
	"testing"
	"time"

	"github.com/jichall/pismo/src/database"
)

func TestInsert(t *testing.T) {

	transaction := Transaction{
		Amount:  -50.64,
		Date:    time.Now(),
		Account: 1,
		Type:    1,
	}

	mock := database.PoolMock{}

	err := transaction.Insert(mock)

	if err != nil {
		t.Fail()
	}
}