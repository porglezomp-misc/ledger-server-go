package main

import (
	"fmt"
)

var currentId int
var entries Entries

func init() {
	RepoCreateEntry(Entry{From: "you", To: "me", Amount: "$20"})
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
