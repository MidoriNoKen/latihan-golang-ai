package config

import (
	"fmt"
	"log"
	"os"

	"github.com/MidoriNoKen/latihan-golang-ai/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Set default values if not specified
	if dbUser == "" {
		dbUser = "root"
	}
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbName == "" {
		dbName = "latihan_golang"
	}

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		dbUser, dbPass, dbHost, dbPort, dbName)

	log.Printf("Connecting to database at %s:%s/%s...", dbHost, dbPort, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Warning: Failed to connect to database: %v", err)
		log.Println("Database connection is not established, but application will continue to run.")
		return
	}

	log.Println("Database connection established successfully.")

	// Run database migrations
	log.Println("Running database migrations...")
	if err := DB.AutoMigrate(&domain.User{}); err != nil {
		log.Printf("Warning: Database migration failed: %v", err)
	} else {
		log.Println("Database migration completed successfully.")
	}
}
