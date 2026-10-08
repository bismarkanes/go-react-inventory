package model

import "time"

type Item struct {
	ID            uint
	Stock         int `gorm:"default:0"`
	ReservedStock int `gorm:"default:0"`
}

type ReservationItem struct {
	ID        uint
	UserID    string
	ItemID    string
	Quantity  int
	ExpiresAt time.Time
	Confirmed bool `gorm:"default:false"`
	ConfirmAt time.Time
}
