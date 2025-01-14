package services

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/sssamuelll/portfolio_backend/models"
	"github.com/sssamuelll/portfolio_backend/storage"
)

// CreateUser inserta un usuario en la base de datos
func CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	_, err := storage.DB.Exec(query, user.Username, user.Email, user.Password)
	return err
}

// GetUserByUsername obtiene un usuario por su username
func GetUserByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, password, email, secret_totp FROM users WHERE username = ?`
	row := storage.DB.QueryRow(query, username)

	var u models.User
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.SecretTOTP)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

// HashPassword hashea la contraseña usando bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) // cost=10
	return string(bytes), err
}

// CheckPassword compara la contraseña en texto plano con el hash de la DB
func CheckPassword(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

// SaveTOTPSecret guarda el secret TOTP en la DB
func SaveTOTPSecret(username, secret string) error {
	query := `UPDATE users SET secret_totp = ? WHERE username = ?`
	_, err := storage.DB.Exec(query, secret, username)
	return err
}

// GetTOTPSecret obtiene el secret TOTP de un usuario
func GetTOTPSecret(username string) (string, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	return user.SecretTOTP, nil
}
