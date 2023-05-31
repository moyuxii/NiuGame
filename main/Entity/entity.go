package Entity

import "gorm.io/gorm"

// 卡牌
type Eumn string

type Card struct {
	Number int
	Decor  Eumn
}
type Cards []Card

type Player struct {
	gorm.Model
	RoomId     string
	PlayerName string
	PlayerRole int `gorm:"default:2"`
}

type Room struct {
	gorm.Model
	RoomId     string `form:"roomId" json:"roomId" `
	RoomPasswd string
	BelongCust string
	Enable     bool `gorm:"default:true"`
}

type Gaming struct {
	gorm.Model
	RoomId string
}

type Customer struct {
	Name   string
	Passwd string
}

type Result struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
