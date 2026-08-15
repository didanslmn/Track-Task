package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Env      string
	Database DBConfig
	Server   ServerConfig
	JWT      JWTConfig
}

type DBConfig struct {
	Host           string
	Port           int
	DBName         string
	User           string
	Password       string
	SSLMode        string
	ChannelBinding string
}

type ServerConfig struct {
	Port         string
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type JWTConfig struct {
	SecretKey      string
	AccessTokenTTL time.Duration
	RefreshToken   time.Duration
	Issuer         string
}

func LoadConfig() (*Config, error) {
	env := getEnv("ENV", "development")

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return nil, fmt.Errorf("DB password wajib diisi (tidak ada default yang aman)")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT secret wajib diisi (tidak ada default yang aman)")
	}

	return &Config{
		Env: env,
		Database: DBConfig{
			Host:           getEnv("DB_HOST", "ep-delicate-cake-aowbpine-pooler.c-2.ap-southeast-1.aws.neon.tech"),
			Port:           getEnvAsInt("DB_PORT", 5432),
			DBName:         getEnv("DB_NAME", "neondb"),
			User:           getEnv("DB_USER", "neondb_owner"),
			Password:       dbPassword,
			SSLMode:        getEnv("DB_SSLMODE", "require"),
			ChannelBinding: getEnv("DB_CHANNEL_BINDING", "require"),
		},
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8080"),
			Host:         getEnv("SERVER_HOST", "localhost"),
			ReadTimeout:  getEnvDuration("SERVER_READ_TIMEOUT", 15*time.Second),
			WriteTimeout: getEnvDuration("SERVER_WRITE_TIMEOUT", 15*time.Second),
		},
		JWT: JWTConfig{
			SecretKey:      jwtSecret,
			AccessTokenTTL: getEnvDuration("JWT_ACCESS_TTL", 15*time.Minute),
			RefreshToken:   getEnvDuration("JWT_REFRESH_TTL", 7*24*time.Hour),
			Issuer:         getEnv("JWT_ISSUER", "team-task-app"),
		},
	}, nil
}

// Helper for get environment variables
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
