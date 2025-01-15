package services

import (
	"time"

	"github.com/pquerna/otp/totp"
)

// GenerateTOTPSecret genera un nuevo secreto TOTP
func GenerateTOTPSecret() string {
	secret, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "PortfolioBackend",
		AccountName: "user@example.com", // Cambiar dinámicamente según el usuario
	})
	return secret.Secret()
}

// GenerateTOTPCode genera un código TOTP basado en un secreto
func GenerateTOTPCode(secret string) string {
	code, _ := totp.GenerateCode(secret, time.Now())
	return code
}

// ValidateTOTP valida un código TOTP con un secreto
func ValidateTOTP(secret, code string) bool {
	return totp.Validate(code, secret)
}
