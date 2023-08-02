package main

import (
	"log"

	"github.com/NooFreeNames/ImageEditor/configs"
	"github.com/NooFreeNames/ImageEditor/internal/server"
)

func main() {
	conf, err := configs.New("./configs/.env")
	if err != nil {
		log.Fatalln(err)
	}
	server.Run(conf)
}
