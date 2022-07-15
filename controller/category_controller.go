package controller

import (
	"github.com/labstack/echo/v4"
	"ecommerce/model"
	"ecommerce/db"
	"net/http"
)


//List all category and subcategory
func GetCategory(c echo.Context) (err error) {
	var category [] model.Category

	db.DB.Preload("SubCategory").Find(&category)

	return c.JSON(http.StatusOK, category)

}