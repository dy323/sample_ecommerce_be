package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	seed "ecommerce/seeder/seed"
)

func main() {
	godotenv.Load("../.env")
	
	db_connect := os.Getenv("DB_HOST")
	db, err := gorm.Open(postgres.Open(db_connect), &gorm.Config{})

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