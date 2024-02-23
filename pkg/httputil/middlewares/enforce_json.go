package middlewares

import (
	"net/http"
	"strings"
)

// EnforceJSON is a middleware that ensures the request's Content-Type is application/json.
// It checks the Content-Type header of incoming requests and if the request method is not GET,
// it will respond with an Unsupported Media Type status if the Content-Type is not application/json.
// This middleware is intended to be used with routes that accept JSON payloads.
func EnforceJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			// Check if the request Content-Type is JSON
			if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
				http.Error(w, "Unsupported Content-Type: application/json required", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
