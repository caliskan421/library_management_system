package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DB  DBConfig
	JWT JWTConfig
	App AppConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret string
	Expiry time.Duration
}

type AppConfig struct {
	Port string
	Env  string
}

func Load() *Config {
	_ = godotenv.Load()

	expiry, err := time.ParseDuration(getEnv("JWT_EXPIRY", "24h"))
	if err != nil {
		expiry = 24 * time.Hour
	}

	return &Config{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "libranet"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "default-secret-change-me"),
			Expiry: expiry,
		},
		App: AppConfig{
			Port: getEnv("APP_PORT", "3000"),
			Env:  getEnv("APP_ENV", "development"),
		},
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
