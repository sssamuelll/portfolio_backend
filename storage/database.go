package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./portfolio.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Crear tabla `posts` si no existe
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		image TEXT,
		name TEXT NOT NULL,
		description TEXT,
		category TEXT NOT NULL,
		tags TEXT,
		media TEXT,
		start_date TEXT,
		end_date TEXT,
		link TEXT
	);`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
