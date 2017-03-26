package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func getRequestToken(r *http.Request) string {
	token := r.FormValue("token")
	if token == "" {
		token = r.Header.Get("X-Token")
	}
	return strings.ToLower(token)
}

func aggregateRequestParams(r *http.Request) map[string]string {
	res := make(map[string]string)

	token := getRequestToken(r)
	if token != "" {
		res["token"] = token
	}

	for key, value := range mux.Vars(r) {
		if value != "" {
			res[strings.ToLower(key)] = value
		}
	}

	if err := r.ParseForm(); err == nil {
		for key, value := range r.Form {
			if len(value) > 0 && value[0] != "" {
				res[strings.ToLower(key)] = value[0]
			}
		}
	}

	for key, value := range r.URL.Query() {
		if len(value) > 0 && value[0] != "" {
			res[strings.ToLower(key)] = value[0]
		}
	}

	return res
}
