package main

import (
	"log"
	"test_project_2/internal/config"
	"test_project_2/internal/db"
	"test_project_2/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	err = db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("Application started successfully")

	router := gin.Default()

	router.POST("/register", handler.Register)
	router.GET("/users", handler.GetUsers)
	router.POST("/login", handler.Login)

	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
