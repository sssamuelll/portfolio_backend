package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pquerna/otp/totp"
	"github.com/sssamuelll/portfolio_backend/config"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT genera un token JWT para un usuario autenticado
func GenerateJWT(username string) (string, error) {
	secretKey := config.AppConfig.SecretKey
	expirationTime := time.Now().Add(24 * time.Hour) // Token válido por 24 horas
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateJWT valida un token JWT
func ValidateJWT(tokenString string) (*Claims, error) {
	secretKey := config.AppConfig.SecretKey
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// GenerateTOTP genera un código QR para 2FA basado en TOTP
func GenerateTOTP(username string) (string, string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "PortfolioBackend",
		AccountName: username,
	})
	if err != nil {
		return "", "", err
	}

	// URL para generar un QR code
	return key.URL(), key.Secret(), nil
}

// ValidateTOTP valida un código TOTP generado por el usuario
func ValidateTOTP(secret string, code string) bool {
	return totp.Validate(code, secret)
}
