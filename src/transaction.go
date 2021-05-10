package main

import "time"

type Transaction struct {
	ID int
	Amount string
	Date time.Time
	Account int
	Type int
}