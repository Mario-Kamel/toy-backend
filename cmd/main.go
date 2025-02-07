package main

import (
	"log"

	"github.com/Mario-Kamel/toy-backend/cmd/api"
	"github.com/Mario-Kamel/toy-backend/db"
)

func main() {

	pgStore, err := db.NewPostgresStore()

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
