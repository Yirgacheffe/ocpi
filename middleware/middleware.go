package middleware

import (
	"log"
	"net/http"
	"time"
)

func WithMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s took %s", r.Method, r.URL, time.Since(began))
	})
}
