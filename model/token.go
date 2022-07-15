package model

import "time"

type Token struct {
	TokenID uint `gorm:"primaryKey;autoIncrement:true;"`
	UserID uint
	AuthToken string `gorm:"type:varchar(100);unique;"`
	Generated_at time.Time
	Expires_at time.Time
}