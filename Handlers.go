package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	entryId, err := strconv.Atoi(vars["entryId"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	entry := RepoFindEntry(entryId)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(entry); err != nil {
		panic(err)
	}
}

func EntryCreate(w http.ResponseWriter, r *http.Request) {
	var entry Entry
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err := json.Unmarshal(body, &entry); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	e := RepoCreateEntry(entry)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}
