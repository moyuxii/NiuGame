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
	PlayerId   string
	RoomId     string
	PlayerName string
	HandCards  Cards
}

type Room struct {
	gorm.Model
	RoomId     string
	RoomPasswd string
	BolongCust string
	Enable     bool `gorm:"default:true"`
}

type Customer struct {
	Name   string
	Passwd string
}

type Result struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
