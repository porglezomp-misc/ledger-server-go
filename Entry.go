package main

import (
	"time"
)

type Entry struct {
	Id     int       `json:"id"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	On     time.Time `json:"on"`
	Amount string    `json:"amount"`
}

type Entries []Entry
