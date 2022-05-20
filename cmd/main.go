package main

import (
	"log"
	"os"

	"github.com/luispfcanales/goweb/cmd/httpserver"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if err := httpserver.Run(port); err != nil {
		log.Fatal(err)
	}
}
