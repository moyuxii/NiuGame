package service

import (
	"NiuGame/main/Entity"
	"log"
)

func GetPlayersByRoomId(RoomId string) ([]Entity.Player, int64) {
	PlayerList := []Entity.Player{}
	var count int64
	if err := conn.Where("room_id = ?", RoomId).Find(&PlayerList).Count(&count).Error; err != nil {
		log.Println(err.Error())
		return nil, 9
	}
	return PlayerList, count
}
