package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Global DB variable that will be shared across the app
var DB *gorm.DB

func InitDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Fetch environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	data_src_name := os.Getenv("DB_NAME")

	// Construct the Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, data_src_name)

	// Open connection to the database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Use singular table names for all models
		},
	})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Print("Successfully connected to the database")
}
