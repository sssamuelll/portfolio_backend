package models

type User struct {
	ID         int    `json:"id" db:"id"`
	Username   string `json:"username" db:"username"`
	Password   string `json:"-" db:"password"` // con "-" para no exponerlo en JSON
	Email      string `json:"email" db:"email"`
	SecretTOTP string `json:"-" db:"secret_totp"` // para guardar el secret de TOTP
}
