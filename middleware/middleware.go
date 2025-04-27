package middleware

import (
	"log"
	"net/http"
)

// EnableCORS sets the CORS headers and handles preflight OPTIONS requests
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log entry into the middleware
		log.Printf("CORS: Middleware ENTRY for %s %s", r.Method, r.URL.RequestURI())

		// Set common CORS headers
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Fallback, but less ideal
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") //  Include Authorization if you use it
		w.Header().Set("Access-Control-Max-Age", "86400")                             // Add Max-Age for caching

		// Log header setting
		log.Printf("CORS: Headers SET for %s %s", r.Method, r.URL.RequestURI())

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			log.Printf("CORS: Handling OPTIONS preflight for %s", r.URL.RequestURI()) // Log OPTIONS handling
			w.WriteHeader(http.StatusOK)
			return // Stop processing for OPTIONS
		}

		// Call next handler
		log.Printf("CORS: Calling next handler for %s %s", r.Method, r.URL.RequestURI()) // Log before calling next
		next.ServeHTTP(w, r)
		log.Printf("CORS: Middleware EXIT after next handler for %s %s", r.Method, r.URL.RequestURI()) // Log after calling next
	})
}
