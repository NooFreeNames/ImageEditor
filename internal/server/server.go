// Package server provides a simple server implementation.
package server

import (
	"log"
	"net/http"

	"github.com/NooFreeNames/ImageEditor/configs"
	hndls "github.com/NooFreeNames/ImageEditor/internal/server/handlers"
	mw "github.com/NooFreeNames/ImageEditor/internal/server/middleware"
)

// Function Run starts the server and listens for incoming HTTP requests.
func Run(conf configs.ConfigI) {
	fs := http.FileServer(http.Dir(conf.GetSiteDir()))
	http.Handle("/", mw.LogRequest(fs))
	http.Handle("/ping", mw.LogRequest(http.HandlerFunc(hndls.PingHandler)))
	http.Handle("/image", mw.LogRequest(http.HandlerFunc(hndls.ImageHandler)))

	addr := conf.GetHost() + ":" + conf.GetPort()

	log.Println("Server is listening at http://" + addr + "/")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
