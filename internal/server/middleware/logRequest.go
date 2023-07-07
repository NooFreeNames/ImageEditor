// Package middleware provides a collection of middleware functions for the
// server package.
package middleware

import (
	"log"
	"net/http"
)

// Middleware function to log HTTP requests.
func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			log.Println("Request:", request.Method, request.URL.Path)
			handler.ServeHTTP(response, request)
		})
}
