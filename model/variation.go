package model

import "github.com/bojanz/currency"

type Variation struct {
	VariationID uint `gorm:"primaryKey;"`
	Name string
	ProductID uint
	Quantity uint
	Price currency.Amount 
}