package seed

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"time"
	 "ecommerce/model"
)

func SeedProductDB(db *gorm.DB) {

	var data [] model.Product

	for i:=0; i<5; i++ {

		data = append(data, model.Product {
			ProductID: uint(i+1), 
			Name: faker.FirstName(), 
			CategoryID: uint(i+1), 
			SubCategoryID: uint(i+1), 
			ProfileID: 1, //1 is merchant, 2 is customer
			RegisterDate: time.Now(), 
			Description: faker.Paragraph(),
		})
	}

	db.Create(&data)
}
