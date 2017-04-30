package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func openFile(name string) (*os.File, error) {
	const filePage string = "logs/samatha_%s.log"
	file, err := os.OpenFile(fmt.Sprintf(filePage, name), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	return file, nil
}

func LogAccess(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

func LogCacheAccess(url string, cacheTimestamp time.Time) {
	log.Printf(
		"samantha-cached | %s\t%s",
		cacheTimestamp,
		url,
	)
}

func LogError(message string, prefix string) {
	file, _ := openFile("error")
	defer file.Close()
	if file != nil {
		logger := log.New(file, fmt.Sprintf("%s ", prefix), 3)
		logger.Printf(
			"%s\t",
			message,
		)
	}

}
