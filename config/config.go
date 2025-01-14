package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	SecretKey  string
	Database   string
	IssuerName string
}

var AppConfig Config

func LoadConfig() {
	// Cargar variables desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using defaults.")
	}

	AppConfig = Config{
		Port:       getEnv("PORT", "8080"),
		SecretKey:  getEnv("SECRET_KEY", "your_secret_key"),
		Database:   getEnv("DATABASE", "./portfolio.db"),
		IssuerName: getEnv("ISSUER_NAME", "PortfolioBackend"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
