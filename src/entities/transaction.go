package entities

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jichall/pismo/src/database"
)

type Transaction struct {
	ID      int       `json:"id"`
	Amount  int       `json:"amount"`
	Date    time.Time `json:"timestamp"`
	Account int       `json:"account_id"`
	Type    int       `json:"operation_type_id"`
}

const (
	PayInFull       = 1
	PayInstallments = 2
	Withdraw        = 3
	Payment         = 4
)

func (t *Transaction) String() string {

	format := "Transaction<id: %d amount: %f account_id: %d type: %d time: %v>"

	return fmt.Sprintf(format, t.ID, t.Amount, t.Account, t.Type, t.Date)
}

func (t *Transaction) Insert(db database.DatabaseFunctions) error {

	// Retrieve the user of the transaction
	a := Account{ID: t.Account}
	err := a.Select(db)

	if err != nil {
		return err
	}

	switch t.Type {
	case PayInFull:
	case PayInstallments:
	case Withdraw:
		t.Amount = -t.Amount
	}

	amount := float64(t.Amount) / 100
	credit := float64(a.Credit) / 100

	computedCredit := credit + amount

	log.Printf("%s\n%s\n%f", t.String(), a.String(), computedCredit)

	if computedCredit <= 0 {
		msg := fmt.Sprintf("account %d has no credit available.", a.ID)
		return errors.New(msg)
	}

	timestamp := time.Now().Format("2006-01-02T15:04:05-0700")

	a.Credit = int(computedCredit * 100)
	a.Update(db)

	err = db.Exec(database.InsertTransaction, t.Amount, timestamp, t.Account,
		t.Type)

	return err
}
