package main

import (
	"context"
	"fmt"
	"time"
)

type Transaction struct {
	ID      int
	Amount  string `json:"amount"`
	Date    time.Time
	Account int `json:"account_id"`
	Type    int `json:"operation_type_id"`
}

func (t *Transaction) String() string {

	format := "Transaction<id: %d amount: %s account_id: %d type: %d time: %v>"

	return fmt.Sprintf(format, t.ID, t.Amount, t.Account, t.Type, t.Date)
}

func (t *Transaction) Insert() error {
	conn, err := pool.Acquire(context.Background())

	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Query(context.Background(), insertTransaction,
		[]interface{}{t.Amount, time.Now(), t.Account, t.Type})

	return err
}
