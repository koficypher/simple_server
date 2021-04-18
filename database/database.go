package database

import (
	"fmt"
	"os"

	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabase - contains configurations for setting up connection to a DB server
// and returns a pointer to the Database object
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Connecting to a database server .....")

	loadEnvVariables()

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}

func loadEnvVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
