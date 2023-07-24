// Package server provides a simple server implementation.
package server

import (
	"log"
	"net/http"

	hndls "github.com/NooFreeNames/ImageEditor/internal/server/handlers"
	mw "github.com/NooFreeNames/ImageEditor/internal/server/middleware"
)

// Function Run starts the server and listens for incoming HTTP requests.
func Run() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", mw.LogRequest(fs))
	http.Handle("/ping", mw.LogRequest(http.HandlerFunc(hndls.PingHandler)))
	http.Handle("/image", mw.LogRequest(http.HandlerFunc(hndls.ImageHandler)))

	log.Println("Server is listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
