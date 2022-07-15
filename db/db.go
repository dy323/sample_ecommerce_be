package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
)

var DB *gorm.DB

// ConnectDB opens a connection to the database
func init() {
	godotenv.Load(".env")
	db_connect := os.Getenv("DB_HOST")

	db, err := gorm.Open(postgres.Open(db_connect), &gorm.Config{})

	if (err != nil) {
		panic("can't connect to database")
	}

	DB = db;
}