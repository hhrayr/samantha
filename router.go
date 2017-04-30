package main

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/hhrayr/samantha/httpCache"
	"github.com/hhrayr/samantha/utils"
)

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	routes := getRoutes()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = utils.LogAccess(handler)
		handler = handlerRecover(handler)
		if route.NoCache != true {
			handler = handlerCache(handler)
		}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(handler)
	}

	return router
}

func handlerRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
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
				NewHttpError(err, request.URL.String()).WriteToResponse(w)
			}
		}()
		h.ServeHTTP(w, request)
	})
}

func handlerCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setResponseHeaders := func(header http.Header) {
			for headerKey, headerValue := range header {
				w.Header()[headerKey] = headerValue
			}
		}

		cache := httpCache.NewHttpCache(r.RequestURI)

		if r.Header.Get("X-Reset-Cache") == "" {
			cachedValue, err := cache.GetValue()
			if err != nil {
				utils.LogError(err.Error(), "cache")
			} else if cachedValue != nil {
				utils.LogCacheAccess(r.RequestURI, cachedValue.Timestamp)
				setResponseHeaders(cachedValue.Header)
				w.Header().Set("Samantha-Cached", cachedValue.Timestamp.String())
				w.Write(cachedValue.Data)
				return
			}
		}

		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, r)
		if rec.Code == http.StatusOK {
			err := cache.SetValue(httpCache.NewHttpCachedDataFromResponseRecorder(rec))
			if err != nil {
				utils.LogError(err.Error(), "cache")
			}
		} else {
			w.WriteHeader(rec.Code)
		}
		setResponseHeaders(rec.Header())
		w.Write(rec.Body.Bytes())
	})
}
