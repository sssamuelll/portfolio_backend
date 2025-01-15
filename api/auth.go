package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sssamuelll/portfolio_backend/config"
	"github.com/sssamuelll/portfolio_backend/models"
	"github.com/sssamuelll/portfolio_backend/services"
	// Añade aquí las importaciones que necesites (bcrypt, etc.)
)

// LoginRequest representa el body esperado para un login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse estructura la respuesta que devolvemos al autenticar
type LoginResponse struct {
	Token string `json:"token"`
}

// Login step 1: Verify credentials and send 2FA code
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := services.GetUserByUsername(req.Username)
	if err != nil || !services.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate and send 2FA code
	code := services.GenerateEmailCode()
	if err := services.SendEmail(user.Email, "Your 2FA Code", "Your code: "+code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	// Save code in the database
	if err := services.SavePendingCode(user.Username, code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "2FA code sent to your email"})
}

// Login step 2: Validate 2FA code
type VerifyCodeRequest struct {
	Username string `json:"username" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

func VerifyCode(c *gin.Context) {
	var req VerifyCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := services.GetUserByUsername(req.Username)
	if err != nil || user.PendingCode != req.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid code"})
		return
	}

	// Clear the pending code
	if err := services.ClearPendingCode(req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear code"})
		return
	}

	// Generate JWT
	token, err := services.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

// RegisterRequest representa el body para registrar un usuario
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"    binding:"required"`
	// Añade más campos que sean necesarios
}

// Register endpoint para crear usuarios en la DB
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Verify email is in the allowed list
	if !config.AppConfig.AllowedEmails[req.Email] {
		c.JSON(http.StatusForbidden, gin.H{"error": "Email not allowed for registration"})
		return
	}

	// Hash the password
	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Create the user in the database
	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// SetupTOTP endpoint para generar una URL/Secret para configurar TOTP
func SetupTOTP(c *gin.Context) {
	// Normalmente necesitas saber qué usuario es el que solicita el TOTP
	// Podrías extraerlo del contexto (JWT claims)
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Generar TOTP
	url, secret, err := services.GenerateTOTP(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate TOTP"})
		return
	}

	// Guardar el secret en la DB (asociado al usuario)
	if err := services.SaveTOTPSecret(username.(string), secret); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save TOTP secret"})
		return
	}

	// Devolver al frontend la URL para generar QR y el secret
	c.JSON(http.StatusOK, gin.H{
		"qr_url": url,
		"secret": secret,
	})
}

// VerifyTOTP endpoint para validar un código TOTP
type TOTPVerifyRequest struct {
	Code string `json:"code" binding:"required"`
}

func VerifyTOTP(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req TOTPVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Obtener el secret del usuario en DB
	secret, err := services.GetTOTPSecret(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve TOTP secret"})
		return
	}

	// Validar TOTP
	if !services.ValidateTOTP(secret, req.Code) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid TOTP code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "TOTP code verified successfully"})
}
