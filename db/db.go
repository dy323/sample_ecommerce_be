package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ecommerce/config"
)

var DB *gorm.DB

var envData = config.EnvConfig()

// ConnectDB opens a connection to the database
func init() {

	db, err := gorm.Open(postgres.Open((*envData).DB_HOST), &gorm.Config{})

	if (err != nil) {
		panic("can't connect to database")
	}

	DB = db;
}