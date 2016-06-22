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
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(entries); err != nil {
		panic(err)
	}
}

func EntryShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryId := vars["entryId"]
	fmt.Fprintln(w, "Entry show:", entryId)
}
