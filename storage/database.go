package storage

import (
	"log"

	"github.com/sssamuelll/portfolio_backend/config"
	"github.com/sssamuelll/portfolio_backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase initializes the database connection and runs migrations
func InitDatabase() {
	var err error

	// Conexión a la base de datos SQLite
	DB, err = gorm.Open(sqlite.Open(config.AppConfig.Database), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ejecutar migraciones
	RunMigrations()
}

// RunMigrations ejecuta las migraciones de la base de datos
func RunMigrations() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Post{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrations completed successfully.")
}
