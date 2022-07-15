package seed

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"time"
	"math/rand"
	"github.com/bojanz/currency"
	"strconv"
	"ecommerce/model"
)

func SeedVariationDB(db *gorm.DB){

	var ProductData [] model.Product
	var VariationData [] model.Variation

	db.Select("ProductID").Find(&ProductData)

	rand.Seed(time.Now().Unix())

	for i:=1; i<=15; i++ {

		//random product id
		r := rand.Int() % len(ProductData)

		//rand quantity
		n := rand.Intn(20 - 10 + 1) + 10

		//rand price
		p := rand.Intn(100.00 - 90.00 + 1.00) + 90.00

		price, _ := currency.NewAmount(strconv.Itoa(p), "USD")

		VariationData = append(VariationData, model.Variation {VariationID: uint(i), Name: faker.Username(), ProductID: ProductData[r].ProductID, Quantity: uint(n), Price: price})
	}

	db.Create(&VariationData)

}
