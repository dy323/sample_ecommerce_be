package model

type Category struct {
	CategoryID uint `gorm:"primaryKey;"`
	Name string
	Descp string
	SubCategory []SubCategory `gorm:"foreignKey:CategoryID;references:CategoryID;"`
}


