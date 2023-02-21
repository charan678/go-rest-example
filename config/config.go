package config

import (
	"os"
	"strconv"
)

type Config struct {
	Host string
	Port int
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		Host: getEnv("HOST", "localhost"),
		Port: getEnvAsInt("PORT", 8080),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
