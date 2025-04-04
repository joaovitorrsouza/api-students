package main

import (
	"log"

	"github.com/joaovitorrsouza/api-students/api"
)

func main() {
	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
