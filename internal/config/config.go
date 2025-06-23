package config

import (
	"os"
)

type Config struct {
    DatabaseURL string
    Port        string
}

func Load() *Config {
    return &Config{
        DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@db:5432/DND?sslmode=disable"),
        Port:        getEnv("PORT", "8080"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}