package main

import (
	"log"

	"github.com/xadhithiyan/videon/cmd/api"
	"github.com/xadhithiyan/videon/db"
)

func main() {

	dbconn, err := db.CreateDbInstance()
	if err != nil {
		log.Fatal(err)
	}
	db.PingDB(dbconn)

	server := api.CreateAPIServer(":8080", dbconn)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
