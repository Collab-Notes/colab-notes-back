package common

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB establihshes the connection to the database using GORM and PostgreSQL
func ConnectDB(cfg *Config) {
	dsn := cfg.DATABASE_URL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Assign the connection to the global variable DB
	DB = db
}
