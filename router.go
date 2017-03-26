package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = hadlerRecover(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func hadlerRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			recover := recover()
			if recover != nil {
				switch t := recover.(type) {
				case string:
					log.Println("string " + t)
					err = errors.New(t)
				case error:
					log.Println("error " + t.Error())
					err = t
				default:
					log.Println("default ")
					err = errors.New("Unknown error")
				}
				httpError := NewHttpError(err)
				httpError.SetRequestParameters(aggregateRequestParams(r))
				httpError.SetRequestUrl(fmt.Sprintf("%s %s", r.Method, r.URL))
				httpError.WriteToResponse(w)
				notifyError(err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func notifyError(err error) {
	//log.Println(err.Error())
}
