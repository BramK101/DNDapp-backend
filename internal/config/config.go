package config

import (
	"os"
)

type Config struct {
	DatabaseURL      string
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	DatabasePort     string
	DatabaseSSL      string
	DatabaseName     string
	ResetDatabase    bool
	UrlPort          string
}

func Load() *Config {
	return &Config{
		DatabaseURL:      getEnv("DATABASE_URL", "postgres://user:password@db:5432/DND?sslmode=disable"),
		DatabaseHost:     getEnv("DATABASE_HOST", "db"),
		DatabaseUser:     getEnv("DATABASE_USER", "DND"),
		DatabasePassword: getEnv("DATABASE_PASSWORD", "password"),
		DatabasePort:     getEnv("DATABASE_PORT", "5432"),
		DatabaseSSL:      getEnv("DATABASE_SSL", "false"),
		DatabaseName:     getEnv("DATABASE_NAME", "db"),
		UrlPort:          getEnv("PORT", ":8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
