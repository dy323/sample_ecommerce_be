package main

import (
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"ecommerce/model"
	"ecommerce/config"
)

var envData = config.EnvConfig()

func migration(db *gorm.DB) {
	
	db.AutoMigrate(model.User{}, model.Profile{}, model.Token{}, model.Product{}, model.Category{}, model.SubCategory{}, model.Variation{}, model.MemberShip{})
}

func main () {
	db, err := gorm.Open(postgres.Open((*envData).DB_HOST), &gorm.Config{})

	dba, err := db.DB()

	if (err != nil) {
		panic("can't connect to database")
	}

	migration(db)

	defer dba.Close()

}