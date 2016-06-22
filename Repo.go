package main

import (
	"fmt"
	"time"
)

var currentId int
var entries Entries

func init() {
	RepoCreateEntry(Entry{
		Name: "Wow Food",
		Actions: []Transaction{
			Transaction{Account: "expenses:food:dining", Amount: "$20.00"},
			Transaction{Account: "assets:bank:checking", Amount: ""},
		},
	})
}

func RepoFindEntry(id int) Entry {
	for _, e := range entries {
		if e.Id == id {
			return e
		}
	}
	return Entry{}
}

func RepoCreateEntry(e Entry) Entry {
	if e.On == nil {
		time := time.Now()
		e.On = &time
	}

	currentId += 1
	e.Id = currentId
	entries = append(entries, e)
	return e
}

func RepoDestroyEntry(id int) error {
	for i, e := range entries {
		if e.Id == id {
			entries = append(entries[:i], entries[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
