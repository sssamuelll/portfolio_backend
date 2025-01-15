package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sssamuelll/portfolio_backend/services"
)

// TestEmailRequest defines the request body for testing email
type TestEmailRequest struct {
	Recipient string `json:"recipient" binding:"required"`
}

// SendTestEmail sends a test email to the specified recipient
func SendTestEmail(c *gin.Context) {
	var req TestEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := services.SendEmail(req.Recipient, "Test Email", "Hello World")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test email sent successfully"})
}
