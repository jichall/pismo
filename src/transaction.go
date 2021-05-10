package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type Transaction struct {
	ID      int       `json:"id"`
	Amount  float64   `json:"amount"`
	Date    time.Time `json:"timestamp"`
	Account int       `json:"account_id"`
	Type    int       `json:"operation_type_id"`
}

func (t *Transaction) String() string {

	format := "Transaction<id: %d amount: %f account_id: %d type: %d time: %v>"

	return fmt.Sprintf(format, t.ID, t.Amount, t.Account, t.Type, t.Date)
}

func (t *Transaction) Insert() error {
	conn, err := pool.Acquire(context.Background())

	if err != nil {
		return err
	}
	defer conn.Release()

	amount := strconv.FormatFloat(t.Amount, 'f', 5, 64)
	timestamp := time.Now().Format("2006-01-02T15:04:05-0700")

	_, err = conn.Exec(context.Background(), insertTransaction,
		amount, timestamp, t.Account, t.Type)

	return err
}
