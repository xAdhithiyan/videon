package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/xadhithiyan/videon/db"
)

func main() {
	dbconn, err := db.CreateDbInstance()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(dbconn, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migration", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		err := m.Up()
		if err == migrate.ErrNoChange {
			log.Print("Databse up to date")
		} else if err != nil {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		err := m.Steps(-1)

		if err == migrate.ErrNoChange {
			log.Print("Databse fully back")
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
