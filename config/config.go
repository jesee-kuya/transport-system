package config

import (
	"fmt"
	"log"
	"time"

	"github.com/jesee-kuya/transport-system/domain"
	"github.com/joho/godotenv"
)

// Load loads configuration from environment variables.
func Load() (*domain.Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}
	environment := GetEnv("ENVIRONMENT", "development")
	config := &domain.Config{
		Environment: environment,
		Server: domain.ServerConfig{
			Host: GetEnv("SERVER_HOST", "0.0.0.0"),
			Port: GetEnvAsInt("SERVER_PORT", 8080),
		},
		Database: domain.DatabaseConfig{
			Host:            GetEnv("DB_HOST", "localhost"),
			Port:            GetEnvAsInt("DB_PORT", 5432),
			User:            GetEnv("DB_USER", "postgres"),
			Password:        GetEnv("DB_PASSWORD", ""),
			DBName:          GetEnv("DB_NAME", "gymbro"),
			SSLMode:         GetEnv("DB_SSLMODE", "disable"),
			MaxOpenConns:    GetEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    GetEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: GetEnvAsDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
		},
		JWT: domain.JWTConfig{
			SecretKey:       GetEnv("JWT_SECRET_KEY", ""),
			TokenExpiration: GetEnvAsDuration("JWT_TOKEN_EXPIRATION", 24*time.Hour),
		},
	}

	// Validate required fields
	if config.JWT.SecretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY is required")
	}
	if config.Database.Password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	return config, nil
}
