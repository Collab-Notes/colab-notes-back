package common

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL string `env:"DATABASE_URL"`
}

var config *Config
var once sync.Once

func LoadConfig() *Config {
	// Load config only once
	once.Do(func() {
		// Load environment variables from .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		// Map environment variables to Config struct
		cfg := Config{}
		err = env.Parse(&cfg)
		if err != nil {
			log.Fatalf("Error parsing env variables: %v", err)
		}

		config = &cfg
	})

	return config
}
