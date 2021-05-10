package main

import (
	"context"
	"fmt"
)

type Account struct {
	ID       int    `json:"id"`
	Document string `json:"document_number"`
}

func (a *Account) String() string {
	return fmt.Sprintf("Account<id: %d document: %s>", a.ID, a.Document)
}

func (a *Account) Insert() error {
	conn, err := pool.Acquire(context.Background())

	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), insertAccount,
		[]interface{}{a.Document})

	return err
}

func (a *Account) Select() error {
	conn, err := pool.Acquire(context.Background())

	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Query(context.Background(), selectAccount,
		[]int{a.ID})

	return err
}

