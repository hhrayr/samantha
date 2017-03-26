package main

import "net/http"

func main() {
	router := newRouter()
	http.ListenAndServe(":4040", router)
}
