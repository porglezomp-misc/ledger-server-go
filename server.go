package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Entry struct {
	From   string    `json:"from"`
	To     string    `json:"to"`
	On     time.Time `json:"on"`
	Amount float64   `json:"amount"`
}

type Entries []Entry

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/entries", EntryIndex)
	router.HandleFunc("/entries/{entryId}", EntryShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func EntryIndex(w http.ResponseWriter, r *http.Request) {
	entries := Entries{
		Entry{From: "me", To: "you"},
	}

	json.NewEncoder(w).Encode(entries)
}

func EntryShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryId := vars["entryId"]
	fmt.Fprintln(w, "Entry show:", entryId)
}
