package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s, %s, %v\n", r.Method, r.URL.Path, time.Now())

		next.ServeHTTP(w, r)
	})

}