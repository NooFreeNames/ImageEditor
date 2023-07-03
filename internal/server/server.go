// Package server provides a simple server implementation.
package server

import (
	"log"
	"net/http"
)

// Middleware function to log HTTP requests
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			log.Println("Request: ", request.Method, request.URL.Path)
			handler.ServeHTTP(response, request)
		})
}

// Function Run starts the server and listens for incoming HTTP requests.
func Run() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", logRequest(fs))
	http.Handle("/ping", logRequest(http.HandlerFunc(pingHandler)))

	log.Println("Server is listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// pingHandler is the handler function for the "/ping" URL. It writes
// the string "PONG" as the response body. Any error that occurs while writing
// the response is logged.
func pingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("PONG"))
	if err != nil {
		log.Println(err)
	}
}
