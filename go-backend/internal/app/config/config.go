package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	SMSAuth SMSAuthConfig
}

type SMSAuthConfig struct {
	CodeLength int
	ExpiresIn int
	Key string
	URL string
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	URL string
}

type JWTConfig struct {
	Secret            string
	RefreshSecret     string
	TTLSeconds        int
	RefreshTTLSeconds int
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			URL: getEnv("DATABASE_URL", "postgres://localhost/go_next_pizza?sslmode=disable"),
		},
		JWT: JWTConfig{
			Secret:           getEnv("JWT_SECRET", "your-secret-key"),
			RefreshSecret: getEnv("JWT_REFRESH_SECRET", "your-refresh-secret-key"),
			TTLSeconds:       getEnvAsInt("JWT_TTL_SECONDS", 3600),
			RefreshTTLSeconds: getEnvAsInt("JWT_REFRESH_TTL_SECONDS", 604800),
		},
		SMSAuth: SMSAuthConfig{
			CodeLength: getEnvAsInt("VERIFICATION_CODE_LENGTH", 6),
			ExpiresIn: getEnvAsInt("VERIFICATION_CODE_EXPIRES_IN", 60),
			Key: getEnv("VERIFICATION_CODE_KEY", "your-api-key"),
			URL: getEnv("VERIFICATION_CODE_URL", "https://api.exolve.ru/messaging/v1/SendSMS"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}