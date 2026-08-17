package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    string
	MongoURI      string
	MongoDatabase string
}

func LoadConfig() *Config {

	envPaths := []string{".env", "../.env", "../../.env"}
	loaded := false
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			loaded = true
			break
		}
	}
	if !loaded {
		log.Println("Warning: .env file not found")
	}

	config := &Config{
		ServerPort:    os.Getenv("SERVER_PORT"),
		MongoURI:      os.Getenv("MONGO_URI"),
		MongoDatabase: os.Getenv("MONGO_DATABASE"),
	}

	if config.ServerPort == "" {
		config.ServerPort = "8080"
	}

	if config.MongoURI == "" {
		log.Fatal("MONGO_URI is not configured")
	}

	if config.MongoDatabase == "" {
		log.Fatal("MONGO_DATABASE is not configured")
	}

	return config
}
