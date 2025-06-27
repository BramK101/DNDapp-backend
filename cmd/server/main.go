package main

import (
	"fmt"
	"log"

	"github.com/BramK101/DNDapp-backend/internal/config"
	"github.com/BramK101/DNDapp-backend/internal/handlers"
	"github.com/BramK101/DNDapp-backend/internal/models"
	"github.com/BramK101/DNDapp-backend/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DatabaseHost,
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseName,
		cfg.DatabasePort,
		cfg.DatabaseSSL,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//TODO: hardcoded. ga dit doen voor alle
	db.AutoMigrate(&models.User{})

	allServices := services.NewServices(db)

	allHandlers := handlers.NewHandlers(allServices)

	// Start server (simplified)
	log.Println("Server starting on ", cfg.UrlPort)
	// userHandler.SetupRoutes()
	allHandlers.SetupRoutes(cfg)
}
