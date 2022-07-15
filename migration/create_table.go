package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"os"
	model "ecommerce/model"
)

func migration(db *gorm.DB) {
	
	db.AutoMigrate(model.User{}, model.Profile{}, model.Token{}, model.Product{}, model.Category{}, model.SubCategory{}, model.Variation{}, model.MemberShip{})
}

func main () {
	godotenv.Load("../.env")

	db_connect := os.Getenv("DB_HOST")
	db, err := gorm.Open(postgres.Open(db_connect), &gorm.Config{})

	dba, err := db.DB()

	if (err != nil) {
		panic("can't connect to database")
	}

	migration(db)

	defer dba.Close()

}