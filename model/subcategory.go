package model

type SubCategory struct {
	SubCategoryID uint `gorm:"primaryKey;"`
	Name string
	Descp string
	CategoryID uint
}