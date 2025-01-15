package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sssamuelll/portfolio_backend/services"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		claims, err := services.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Añadir claims al contexto
		c.Set("username", claims.Username)
		c.Next()
	}
}
