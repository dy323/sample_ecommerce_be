package seed

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	 "ecommerce/model"
)

func SeedMemberShipDB(db *gorm.DB) {

	var data [] model.MemberShip

	for i:=1; i<=2; i++ {

		data = append(data, model.MemberShip {MemberShipID: uint(i), Type: faker.LastName()})
	}

	db.Create(&data)
}

