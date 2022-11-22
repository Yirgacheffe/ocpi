package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func WithMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()
		next.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"method":     r.Method,
			"path":       r.URL,
			"latency_ns": time.Since(began),
		}).Info("request details")

		// log.Printf("%s %s took %s", r.Method, r.URL, time.Since(began))
	})
}
