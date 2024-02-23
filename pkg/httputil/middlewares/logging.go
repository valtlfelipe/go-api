package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Custom ResponseWriter to capture the response status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// LogRequest is a middleware that logs all incoming HTTP requests.
// It wraps the response writer to capture the status code, logs the request details,
// including the remote address, request method, URL path, protocol, status code,
// response time, referer, and user-agent, then passes the request down the middleware chain.
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := responseWriter{ResponseWriter: w}
		defer func() {
			log.Printf("%s - - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\"\n",
				r.RemoteAddr,
				start.Format("02/Jan/2006:15:04:05 -0700"),
				r.Method,
				r.URL.Path,
				r.Proto,
				rw.statusCode,
				time.Since(start)/time.Millisecond,
				r.Header.Get("Referer"),
				r.Header.Get("User-Agent"),
			)
		}()

		next.ServeHTTP(&rw, r)
	})
}
