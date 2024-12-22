package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Holds configuration values for database connection.
type Config struct {
	Host				string
	Port				string
	User				string
	Password		string
	DBName			string
	SSLMode			string
}

// Global variable storing the instance of the database connection.
var DB *gorm.DB

// Initialize database connection.
func InitDB(cfg Config) {

	// Generate DSN (Data Source Name) from configuration values.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	// Open connectionb to PostgreSQL database.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the `User` model.
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database.")

	DB = db
}
