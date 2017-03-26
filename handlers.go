package main

import (
	"fmt"
	"net/http"

	"github.com/hhrayr/samantha/configs"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Welcome to samantha!\nFrom: %s", configs.GetEnv())))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	jw := newJsonWriter(aggregateRequestParams(r), w)
	jw.setRequestMethod(r.Method)
	jw.writeApiMethodInvokeResult()
}
