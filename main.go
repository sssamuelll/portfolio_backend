package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/portfolio-backend/api"
	"github.com/portfolio-backend/middlewares"
	"github.com/portfolio-backend/storage"
)

func main() {
	storage.InitDatabase()

	router := gin.Default()

	// Rutas públicas
	router.GET("/api/public/posts", api.GetPublicPosts)

	// Rutas privadas
	private := router.Group("/api/private")
	private.Use(middlewares.AuthenticateJWT())
	{
		private.POST("/posts", api.CreatePost)
	}

	log.Println("Server is running on http://localhost:8080")
	router.Run(":8080")
}
