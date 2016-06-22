package main

import (
	"time"
)

type Entry struct {
	From   string    `json:"from"`
	To     string    `json:"to"`
	On     time.Time `json:"on"`
	Amount float64   `json:"amount"`
}

type Entries []Entry
