package model

import (
	"time"
)

type Product struct {
	ProductID uint `gorm:"primaryKey;"`
	Name string
	CategoryID uint `gorm:"unique;"`
	SubCategoryID uint `gorm:"unique;"`
	ProfileID uint
	RegisterDate time.Time 
	Description string
	Category Category `gorm:"foreignKey:CategoryID;references:CategoryID;"`
	SubCategory SubCategory `gorm:"foreignKey:SubCategoryID;references:SubCategoryID;"`
	Variation []Variation `gorm:"foreignKey:ProductID;references:ProductID;"`
}