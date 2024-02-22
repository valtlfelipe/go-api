package config

import (
	"os"
)

type Config struct {
	RedisURL string
	Port     string
}

func NewConfig() *Config {
	return &Config{
		RedisURL: getEnv("REDIS_URL", ""),
		Port:     getEnv("PORT", "localhost:8090"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
