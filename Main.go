package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/entries", EntryIndex)
	router.HandleFunc("/entries/{entryId}", EntryShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
