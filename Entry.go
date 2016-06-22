package main

import (
	"time"
)

type Entry struct {
	Id      int           `json:"id"`
	On      *time.Time    `json:"on"`
	Name    string        `json:"name"`
	Actions []Transaction `json:"actions"`
}

type Transaction struct {
	Account string `json:"account"`
	Amount  string `json:"amount"`
}

type Entries []Entry
