package main

import (
	"net/http"
)

type Route struct {
	Method       string
	Pattern      string
	NoCache      bool
	AuthRequired bool
	HandlerFunc  http.HandlerFunc
}

type Routes []*Route

func getRoutes() Routes {
	return Routes{
		&Route{
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: indexHandler,
			NoCache:     true,
		},
		&Route{
			Method:      "GET",
			Pattern:     "/health",
			HandlerFunc: healthHandler,
			NoCache:     true,
		},
		&Route{
			Method:      "GET",
			Pattern:     "/api/{apidomain}/{apimethod}",
			HandlerFunc: apiHandler,
			NoCache:     true,
		},
		&Route{
			Method:      "POST",
			Pattern:     "/api/{apidomain}/{apimethod}",
			HandlerFunc: apiHandler,
			NoCache:     true,
		},
	}
}
