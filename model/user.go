package model

import "time"

type User struct {
	UserID uint `gorm:"primaryKey;"`
	Username string
	Email string
	Password string
	Date time.Time
	MemberShip uint
	Profile Profile `gorm:"foreignKey:UserID;references:UserID;"`
	Token Token `gorm:"foreignKey:UserID;references:UserID;"`
}