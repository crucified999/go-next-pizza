package apiserver

import (
	"log/slog"
	"os"
	"strconv"
)


type Config struct {
	BindAddr string
	LogLevel slog.Level
	DataBaseUrl string 
	JWTSecret string
	JWTTTLSeconds int
	RefreshTTLSeconds int
}

func NewConfig() *Config {
	jwtTTL, _ := strconv.Atoi(os.Getenv("JWT_TTL"))
	refreshTTL, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TTL"))

    if jwtTTL == 0 {
        jwtTTL = 900
    }
    if refreshTTL == 0 {
        refreshTTL = 60 * 60 * 24 * 7
    }

	return &Config{
		BindAddr: os.Getenv("BIND_ADDR"),
		LogLevel: slog.LevelDebug,
		DataBaseUrl: os.Getenv("DATABASE_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTTTLSeconds: jwtTTL,
		RefreshTTLSeconds: refreshTTL,
	}
}