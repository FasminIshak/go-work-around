package db

import (
	"database/sql"
	"fmt"
	"log"
	"test_project_2/internal/models"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB(cfg *models.Config) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to database")

	// _, err = DB.Exec(`
	// CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	username VARCHAR(50) UNIQUE NOT NULL,
	// 	password VARCHAR(255) NOT NULL,
	// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	// )`)
	// if err != nil {
	// 	return fmt.Errorf("error creating users table: %v", err)
	// }

	// log.Println("Successfully created tables")
	return nil
}
