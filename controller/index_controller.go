package controller 

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"math/rand"
	"time"
	"ecommerce/model"
	"ecommerce/db"
)

type num [] int

func (list num) contains(element int) bool {
	for _, v := range list {
		if element == v {
			return true
		}
	}
	return false
}

func retrieve(db *gorm.DB, number *num, ch chan interface{}) {
	var data [] model.Product
	db.Limit(len(*number)).Order("product_id desc").Preload("Variation").Preload("Category").Preload("SubCategory").Find(&data)
	ch <- data
}

//Product list which commonly see in homepage
func GetIndex(c echo.Context) (err error) {

	var feature [] model.Product

	ch := make(chan interface{})

	var number num 

	collection := make(map[string]interface{})

	rand.Seed(time.Now().Unix())

	defer close(ch)

	for i:=0; i<15; i++ {
		r := rand.Int() % 10
		if number.contains(r) == false {
			number = append(number, r)
		}
	}

	db.DB.Preload("Variation").Preload("Category").Preload("SubCategory").Find(&feature, number)

	go retrieve(db.DB, &number, ch)

	data := <-ch 

	collection["feature"] = feature
	collection["latest"] = data

	return c.JSON(http.StatusOK, collection)
}