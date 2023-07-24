package handlers

import (
	"log"
	"net/http"
)

// PingHandler is the handler function for the "/ping" URL. It writes
// the string "PONG" as the response body. Any error that occurs while writing
// the response is logged.
func PingHandler(response http.ResponseWriter, request *http.Request) {
	_, err := response.Write([]byte("PONG"))
	if err != nil {
		log.Println("Error writing response: ", err)
	}
}
