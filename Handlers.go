package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func EntryIndex(w http.ResponseWriter, r *http.Request) {
	entries := Entries{
		Entry{From: "me", To: "you"},
		Entry{From: "you", To: "me"},
	}

	json.NewEncoder(w).Encode(entries)
}

func EntryShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryId := vars["entryId"]
	fmt.Fprintln(w, "Entry show:", entryId)
}
