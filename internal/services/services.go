package services

import "github.com/BramK101/DNDapp-backend/internal/database"

type Services struct {
    db *database.DB
}

func NewServices(db *database.DB) *Services {
	return &Services{db: db}
}