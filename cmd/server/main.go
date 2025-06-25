package main

import (
	"log"

	"github.com/BramK101/DNDapp-backend/internal/config"
	"github.com/BramK101/DNDapp-backend/internal/database"
	"github.com/BramK101/DNDapp-backend/internal/handlers"
	"github.com/BramK101/DNDapp-backend/internal/services"
)

func main() {
    // Load configuration
    cfg := config.Load()
    
    // Initialize database
    db, err := database.NewConnection(cfg.DatabaseURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    allServices := services.NewServices(db)
    
    allHandlers := handlers.NewHandlers(allServices)

    if err := db.CreateTables(cfg); err != nil {
        log.Fatal("Failed to create tables:", err)
    }
    
    // Start server (simplified)
    log.Println("Server starting on :8080")
    // userHandler.SetupRoutes()
    allHandlers.SetupRoutes(cfg)
}