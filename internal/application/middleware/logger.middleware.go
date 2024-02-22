package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("Method: %s, URL: %s, Elapsed Time: %s, Origin: %s",
			r.Method, r.URL.String(), time.Since(start), r.Header.Get("Origin"))
	})
}
