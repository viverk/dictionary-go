package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs information about incoming requests.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the start time of the request
		start := time.Now()

		// Serve the request to the next middleware or handler in the chain
		next.ServeHTTP(w, r)

		// Record the end time of the request
		end := time.Now()

		// Calculate the duration of the request
		duration := end.Sub(start)

		// Log the request information
		log.Printf("%s - %s %s %s %v", r.RemoteAddr, r.Method, r.RequestURI, r.Proto, duration)
	})
}
