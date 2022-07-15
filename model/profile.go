package model

type Profile struct {
	ProfileID uint `gorm:"primaryKey;"`
	UUID string `gorm:"type:varchar(100);unique;"`
	UserID uint
	Name string
	Street string
	PostCode string
	State string
	Country string
	Telephone uint
	Product []Product `gorm:"foreignKey:ProfileID;references:ProfileID;"`
}