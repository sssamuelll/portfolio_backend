package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/sssamuelll/portfolio_backend/models"
	"github.com/sssamuelll/portfolio_backend/storage"
)

// CreateUser inserta un usuario en la base de datos
func CreateUser(user *models.User) error {
	return storage.DB.Create(user).Error
}

// GetUserByUsername obtiene un usuario por su username
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := storage.DB.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// HashPassword hashea la contraseña usando bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPassword compara la contraseña en texto plano con el hash de la DB
func CheckPassword(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

// SaveTOTPSecret guarda el secret TOTP en la DB
func SaveTOTPSecret(username, secret string) error {
	return storage.DB.Model(&models.User{}).
		Where("username = ?", username).
		Update("secret_totp", secret).Error
}

// GetTOTPSecret obtiene el secret TOTP de un usuario
func GetTOTPSecret(username string) (string, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	return user.SecretTOTP, nil
}

// SavePendingCode saves the 2FA code for a user
func SavePendingCode(username, code string) error {
	return storage.DB.Model(&models.User{}).
		Where("username = ?", username).
		Update("pending_code", code).Error
}

// ClearPendingCode clears the 2FA code for a user
func ClearPendingCode(username string) error {
	return storage.DB.Model(&models.User{}).
		Where("username = ?", username).
		Update("pending_code", "").Error
}
