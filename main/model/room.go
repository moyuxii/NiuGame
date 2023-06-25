package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomId     string `form:"room_Id" json:"roomId" `
	RoomPasswd string `json:"room_passwd"`
	BelongCust string `json:"belong_cust"`
	Enable     bool   `gorm:"default:true" json:"enable"`
}
