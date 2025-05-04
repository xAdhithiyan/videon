package main

import (
	"log"

	"github.com/xadhithiyan/videon/cmd/api"
)

func main() {
	server := api.CreateAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
