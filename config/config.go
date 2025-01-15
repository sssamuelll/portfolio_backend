package config

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	SecretKey      string
	Database       string
	IssuerName     string
	AllowedEmails  map[string]bool
	EmailSender    string
	EmailPassword  string
	SMTPServer     string
	SMTPPort       string
	DKIMPrivateKey string
	DKIMDomain     string
	DKIMSelector   string
}

var AppConfig Config

func LoadConfig() {
	// Cargar variables desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using defaults.")
	}

	// Cargar la clave privada DKIM desde el archivo
	dkimKeyPath := getEnv("DKIM_PRIVATE_KEY_PATH", "")
	dkimPrivateKey, err := ioutil.ReadFile(dkimKeyPath)
	if err != nil {
		log.Fatalf("Failed to load DKIM private key: %v", err)
	}

	AppConfig = Config{
		Port:           getEnv("PORT", "8080"),
		SecretKey:      getEnv("SECRET_KEY", "your_secret_key"),
		Database:       getEnv("DATABASE", "./portfolio.db"),
		IssuerName:     getEnv("ISSUER_NAME", "PortfolioBackend"),
		AllowedEmails:  parseAllowedEmails(getEnv("ALLOWED_EMAILS", "")),
		EmailSender:    getEnv("EMAIL_SENDER", ""),
		EmailPassword:  getEnv("EMAIL_PASSWORD", ""),
		SMTPServer:     getEnv("SMTP_SERVER", ""),
		SMTPPort:       getEnv("SMTP_PORT", "587"),
		DKIMPrivateKey: string(dkimPrivateKey),
		DKIMDomain:     getEnv("DKIM_DOMAIN", "example.com"),
		DKIMSelector:   getEnv("DKIM_SELECTOR", "default"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func parseAllowedEmails(emails string) map[string]bool {
	emailList := strings.Split(emails, ",")
	emailMap := make(map[string]bool, len(emailList))
	for _, email := range emailList {
		emailMap[email] = true
	}
	return emailMap
}
