package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	RoomId     string
	PlayerName string
	Lock       bool `gorm:"default:false"`
	PlayerRole int  `gorm:"default:2"`
}
