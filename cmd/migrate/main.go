package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		"postgres://postgres:dsa@127.0.0.1:5432/rest_api?sslmode=disable",
	)

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		if err := m.Up(); err != nil {
			if err == migrate.ErrNoChange {
				log.Println("No changes to apply.")
			} else {
				log.Fatalf("Migration failed: %v", err)
			}
		} else {
			log.Println("Migrate UP Success")
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err == migrate.ErrNoChange {
			log.Println("Migrate down err", err)
		} else {
			log.Println("Migrate Down Success")
		}
	}

}
