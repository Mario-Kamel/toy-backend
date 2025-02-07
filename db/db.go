package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mario-Kamel/toy-backend/config"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	log.Printf("User: %s, Password: %s, DBName: %s", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName)
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (pgStore *PostgresStore) InitStorage() {
	err := pgStore.db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB successfully connected!")
}
