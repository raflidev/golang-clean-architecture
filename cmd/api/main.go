package main

import (
	"log"

	"github.com/raflidev/golang-clean-architecture/internal/server"
	"github.com/raflidev/golang-clean-architecture/pkg/db/postgres"
)

func main() {
	db, err := postgres.NewPostgreeDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	s := server.NewServer(":4000", db)

	if err := s.Run(); err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}
}
