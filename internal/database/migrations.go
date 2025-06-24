package database

import (
	"log"

	"github.com/BramK101/DNDapp-backend/internal/config"
)

func (db *DB) CreateTables(cfg *config.Config) error {

	if cfg.ResetDatabase {
        log.Print("Resetting database.")
		query := `
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;`

		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) NOT NULL,
        password VARCHAR(255),
        email VARCHAR(255) UNIQUE,
		google_id VARCHAR(255) UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}