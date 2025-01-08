package config

import (
	"log"
	"os"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Host                       string `env:"HTTP_HOST" validate:"required"`
	Port                       string `env:"HTTP_PORT" validate:"required,numeric"`
	DatabaseHost               string `env:"DATABASE_HOST" validate:"required"`
	DatabasePort               string `env:"DATABASE_PORT" validate:"required,numeric"`
	DatabaseUser               string `env:"DATABASE_USER" validate:"required"`
	DatabasePassword           string `env:"DATABASE_PASSWORD" validate:"required"`
	DatabaseName               string `env:"DATABASE_NAME" validate:"required"`
	JwtSecret                  string `env:"JWT_SECRET" validate:"required"`
	JwtAccessTokenExpiryMinute string `env:"JWT_ACCESS_TOKEN_EXPIRY_MINUTE" validate:"required,numeric"`
}

var (
	instance *Config
	once     sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		// Load .env file
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment variables")
		}

		cfg := &Config{
			Host:                       getEnv("HTTP_HOST", "localhost"),
			Port:                       getEnv("HTTP_PORT", "8080"),
			DatabaseHost:               getEnv("DATABASE_HOST", "localhost"),
			DatabasePort:               getEnv("DATABASE_PORT", "5432"),
			DatabaseUser:               getEnv("DATABASE_USER", "postgres"),
			DatabasePassword:           getEnv("DATABASE_PASSWORD", "password"),
			DatabaseName:               getEnv("DATABASE_NAME", "go-wish-list"),
			JwtSecret:                  getEnv("JWT_SECRET", ""),
			JwtAccessTokenExpiryMinute: getEnv("JWT_ACCESS_TOKEN_EXPIRY_MINUTE", ""),
		}

		// Validate configuration
		validate := validator.New()
		if err := validate.Struct(cfg); err != nil {
			log.Fatalf("Configuration validation failed: %v", err)
		}

		instance = cfg
	})

	return instance
}

// GetConfig provides access to the loaded configuration.
func GetConfig() *Config {
	if instance == nil {
		log.Fatal("Config is not loaded. Call LoadConfig first.")
	}
	return instance
}

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
