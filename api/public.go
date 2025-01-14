package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sssamuelll/portfolio_backend/services"
)

func GetPublicPosts(c *gin.Context) {
	posts, err := services.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}
