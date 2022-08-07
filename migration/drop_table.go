package main

import (
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"ecommerce/model"
	"ecommerce/config"
)

var envData = config.EnvConfig()

func main () {
	db, err := gorm.Open(postgres.Open((*envData).DB_HOST), &gorm.Config{})

	dba, err := db.DB()

	if (err != nil) {
		panic("can't connect to database")
	}

	defer dba.Close()

	db.Migrator().DropTable(&model.Category{}, &model.SubCategory{}, &model.User{}, &model.MemberShip{}, &model.Profile{}, &model.Product{}, &model.Variation{}, &model.Token{})

}