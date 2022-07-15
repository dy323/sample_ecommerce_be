package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
	 "ecommerce/model"
)

func SeedProfileDB(db *gorm.DB) {

	var data [] model.Profile

	var fetchUser [] model.User

	db.Select("UserID").Find(&fetchUser)

	for i:=0; i<=1; i++ {
		uid := uuid.New()

		data = append(data, model.Profile {ProfileID: uint(i+1), UUID: uid.String(), UserID: fetchUser[i].UserID, Name: faker.Name(), Street: faker.Word(), PostCode: "52100", State: faker.FirstNameMale(), Country: faker.Name(), Telephone: 60178715976})
	}

	db.Create(&data)
}

