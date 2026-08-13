package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WebServerPort string
	WeatherAPIKey string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	return &Config{
		WebServerPort: getServerPort(),
		WeatherAPIKey: getEnv("WEATHER_API_KEY", ""),
	}, nil
}

func getServerPort() string {
	cloudRunPort := os.Getenv("PORT")
	if cloudRunPort != "" {
		return cloudRunPort
	}

	return getEnv("WEB_SERVER_PORT", "8080")
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
