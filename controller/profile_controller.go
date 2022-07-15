package controller

import (
	"ecommerce/db"
	"github.com/labstack/echo/v4"
	"net/http"
	model "ecommerce/model"
)

type figure struct {
	UUID string `json: "uuid"`
}

//See Profile and check their products
func GetProfile(c echo.Context)(err error) {

	var query figure

	var identity model.Profile

	if err := c.Bind(&query); err !=nil {
		return err
	}

	db.DB.Where("uuid = ?", query.UUID).Preload("Product").Find(&identity)

	return c.JSON(http.StatusOK, &identity)
}