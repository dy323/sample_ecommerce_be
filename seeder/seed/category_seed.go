package seed

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	 "ecommerce/model"
)

func SeedCategoryDB(db *gorm.DB) {

	var data [] model.Category

	for i:=0; i<5; i++ {

		data = append(data, model.Category{CategoryID: uint(i+1), Name: faker.LastName(), Descp: faker.Sentence()})
	}

	db.Create(&data)
}

