package model

import "gorm.io/gorm"

type Gaming struct {
	gorm.Model
	RoomId     string `json:"room_id"`
	PlayerName string `json:"player_name"`
	Item1      string `json:"item_1"`
	Item2      string `json:"item_2"`
	Item3      string `json:"item_3"`
	Item4      string `json:"item_4"`
	Item5      string `json:"item_5"`
}
