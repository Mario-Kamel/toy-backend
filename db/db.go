package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mario-Kamel/toy-backend/config"
	_ "github.com/lib/pq"
)

func NewPostgresStorage() (*sql.DB, error) {
	log.Printf("User: %s, Password: %s, DBName: %s", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName)
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
