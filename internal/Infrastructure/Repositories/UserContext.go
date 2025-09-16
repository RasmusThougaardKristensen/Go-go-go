package Repositories

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	connectionString := "user=postgres dbname=User.Management sslmode=disable host=localhost"
	database, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := database.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	return database, nil
}
