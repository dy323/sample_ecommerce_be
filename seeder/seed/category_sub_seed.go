package seed

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"ecommerce/model"
)

func SeedCategorySubDB(db *gorm.DB) {

	var data [] model.SubCategory

	for i:=0; i<5; i++ {

		data = append(data, model.SubCategory{SubCategoryID: uint(i+1), Name: faker.Username(), Descp: faker.Sentence(), CategoryID: uint(i+1)})

	}

	db.Create(&data)
}


