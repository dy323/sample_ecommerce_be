package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"os"
	model "ecommerce/model"
)

func main () {
	godotenv.Load("../.env")
	db_connect := os.Getenv("DB_HOST")
	db, err := gorm.Open(postgres.Open(db_connect), &gorm.Config{})

	dba, err := db.DB()

	if (err != nil) {
		panic("can't connect to database")
	}

	defer dba.Close()

	db.Migrator().DropTable(&model.Category{}, &model.SubCategory{}, &model.User{}, &model.MemberShip{}, &model.Profile{}, &model.Product{}, &model.Variation{}, &model.Token{})

}