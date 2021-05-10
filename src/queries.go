package main

const (
	insertAccount = `INSERT INTO TB_ACCOUNT(ACCOUNT_DOCUMENT) VALUES ($1)`
	selectAccount = `SELECT * FROM TB_ACCOUNT WHERE ACCOUNT_ID = $1`

	insertTransaction = `
	INSERT INTO
		TB_TRANSACTION (
			TRANSACTION_AMOUNT,
			TRANSACTION_DATE,
			ACCOUNT_ID,
			TRANSACTION_TYPE
		)

		VALUES (
			$1, $2, $3, $4
		)
	`
)
