package model

type MemberShip struct {
	MemberShipID uint `gorm:"primaryKey;"`
	Type string
}