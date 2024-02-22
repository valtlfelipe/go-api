package config

import "os"

type Config struct {
	RedisURL string
}

func NewConfig() *Config {
	return &Config{
		RedisURL: getEnv("REDIS_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
