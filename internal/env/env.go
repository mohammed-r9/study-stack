package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config = &envConfig{}

func LoadEnv(path string) {
	if err := godotenv.Load(path); err != nil {

		log.Printf("No .env file found: %v\n", err)
	}

	Config = &envConfig{
		POSTGRES_CONNECTION: mustGetEnv("POSTGRES_CONNECTION"),

		JWT_KEY: []byte(mustGetEnv("JWT_KEY")),
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("error reading environment variable: %v\n", key)
	}
	return value
}

type envConfig struct {
	POSTGRES_CONNECTION string
	// App
	JWT_KEY []byte
}
