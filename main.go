package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sssamuelll/portfolio_backend/api"
	"github.com/sssamuelll/portfolio_backend/config"
	"github.com/sssamuelll/portfolio_backend/middlewares"
	"github.com/sssamuelll/portfolio_backend/storage"
)

func main() {
	// Cargar configuración
	config.LoadConfig()

	// Inicializar la base de datos
	storage.InitDatabase()

	router := gin.Default()
	router.Use(cors.Default())

	// Rutas públicas
	router.POST("/api/public/login", api.Login)
	router.POST("/api/public/register", api.Register)
	router.GET("/api/public/posts", api.GetPublicPosts)
	router.POST("/api/public/verify_code", api.VerifyCode)

	// Rutas privadas
	private := router.Group("/api/private")
	private.Use(middlewares.AuthenticateJWT())
	{
		// Rutas relacionadas con TOTP (para usuarios ya logueados con JWT)
		private.POST("/totp/setup", api.SetupTOTP)
		private.POST("/totp/verify", api.VerifyTOTP)

		// Otros endpoints privados
		private.POST("/posts", api.CreatePost)
	}

	log.Printf("Server is running on http://localhost:%s", config.AppConfig.Port)
	router.Run(":" + config.AppConfig.Port)
}
