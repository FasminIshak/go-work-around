package db

import (
	"database/sql"
	"fmt"
	"test_project_2/internal/models"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(cfg *models.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// Connect to database
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	// Run migrations
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		return fmt.Errorf("error creating migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %v", err)
	}

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

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error running migrations: %v", err)
	}

	// Initialize GORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error initializing GORM: %v", err)
	}

	DB = gormDB
	return nil
}
