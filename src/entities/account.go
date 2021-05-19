package entities

import (
	"fmt"
	"log"

	"github.com/jichall/pismo/src/database"
)

type Account struct {
	ID       int    `json:"id"`
	Document string `json:"document_number"`
	Credit   int    `json:"credit"`
}

func (a *Account) String() string {
	return fmt.Sprintf("Account<id: %d document: %s credit: %d>", a.ID,
		a.Document, a.Credit)
}

func (a *Account) Insert(db database.DatabaseFunctions) error {
	err := db.Exec(database.InsertAccount, a.Document, a.Credit)

	return err
}

func (a *Account) Update(db database.DatabaseFunctions) error {
	return db.Exec(database.UpdateAccount, a.Credit, a.ID)
}

func (a *Account) Select(db database.DatabaseFunctions) error {
	rows, err := db.Query(database.SelectAccount, a.ID)

	for rows.Next() {
		rows.Scan(&a.ID, &a.Document, &a.Credit)
	}

	log.Printf("%s", a.String())

	return err
}
