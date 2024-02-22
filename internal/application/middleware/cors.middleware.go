package middleware

import (
	"net/http"
	"slices"
	"strings"
)

var originAllowlist = []string{
	"http://localhost:5173",
}

var methodAllowlist = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if slices.Contains(originAllowlist, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Add("Vary", "Origin")

			if r.Method == "OPTIONS" && r.Header.Get("Origin") != "" && r.Header.Get("Access-Control-Request-Method") != "" {
				method := r.Header.Get("Access-Control-Request-Method")
				if slices.Contains(methodAllowlist, method) {
					w.Header().Set("Access-Control-Allow-Methods", strings.Join(methodAllowlist, ", "))
					return
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
