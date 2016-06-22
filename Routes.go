package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"EntryIndex",
		"GET",
		"/entries",
		EntryIndex,
	},
	Route{
		"EntryShow",
		"GET",
		"/entries/{entryId}",
		EntryShow,
	},
	Route{
		"EntryCreate",
		"POST",
		"/entries",
		EntryCreate,
	},
}
