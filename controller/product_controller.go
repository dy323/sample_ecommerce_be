package controller 

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"database/sql"
	"strconv"
	"ecommerce/model"
	"ecommerce/db"
)

type criteria struct {
	Term string `json: "term"`
	Min int `json: "min"`
	Max int `json:"max"`
	SubCategory uint `json:"subcategory"`
	Sort string `json:"sort"`
}

//search product detail
func GetProduct(c echo.Context)(err error){

	var result model.Product

	item := c.Param("id")

	db.DB.Where("product_id = ?", item).Preload("Variation").Preload("Category").Preload("SubCategory").Find(&result)

	return c.JSON(http.StatusOK, &result)

}

//filter product
func GetSearch(c echo.Context)(err error) {

	var query criteria

	var result [] model.Product

	if err := c.Bind(&query); err != nil {
		return err
	}

	//fetch product and variations
	db.DB.Raw("SELECT * FROM products WHERE LOWER(Name) LIKE LOWER(@Term) AND (sub_category_id = @Category OR @Category = 0) ORDER BY CASE WHEN @Order = 'ASC' THEN name END ASC, CASE WHEN @Order = 'DSC' THEN name End DESC",
	sql.Named("Term", query.Term), sql.Named("Category", query.SubCategory), sql.Named("Order", query.Sort)).Preload("Variation").Limit(5).Find(&result)

	for n,i := range result {

		//Find whether any variation is within the price range
		pass := false

		for _, k := range i.Variation {
			p,_ := strconv.Atoi(k.Price.Number())
			
			if (p >= query.Min && p <= query.Max) {
				pass = true
			}

		}

		//Else remove the parent
		if (!pass) {
			result = append(result[:n], result[n+1:]...)
		}

	}


	return c.JSON(http.StatusOK, result)

}