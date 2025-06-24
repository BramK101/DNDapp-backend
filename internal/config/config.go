package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
    DatabaseURL string
    ResetDatabase bool
    Port        string
}

func Load() *Config {
    resetDatabse, err := strconv.ParseBool(getEnv("RESET_DATABASE", "true"))
    if err != nil {
        log.Fatal("RESET_DATABSE is not a boolean:", err)
    }

    return &Config{
        DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@db:5432/DND?sslmode=disable"),
        Port:        getEnv("PORT", "8080"),
        ResetDatabase: resetDatabse,
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}