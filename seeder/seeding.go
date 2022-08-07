package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ecommerce/seeder/seed"
	"ecommerce/config"
)

var envData = config.EnvConfig()

func main() {
	db, err := gorm.Open(postgres.Open((*envData).DB_HOST), &gorm.Config{})

	dba, err := db.DB()

	if (err != nil) {
		panic("can't connect to database")
	}

	defer dba.Close()

	seed.SeedMemberShipDB(db)

	seed.SeedUserDB(db)

	seed.SeedProfileDB(db)

	seed.SeedProductDB(db)

	seed.SeedCategoryDB(db)

	seed.SeedCategorySubDB(db)

	seed.SeedVariationDB(db)

}