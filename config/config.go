package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

var AppConfig *Config

func LoadConfig() Config {
	EnvPath := os.Getenv("DOTENV_CONFIG_PATH")
	if EnvPath == "" {
		EnvPath = ".env"
	}
	if err := godotenv.Load(EnvPath); err != nil {
		panic("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	AppConfig = &Config{
		Port: port,
		Env:  os.Getenv("ENV"),
	}
	return *AppConfig
}
