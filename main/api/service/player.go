package service

import (
	"NiuGame/main/Entity"
	"log"
)

func AddPlayer(RoomId string, CustName string, Role int) Entity.Result {
	var newPlayer Entity.Player
	if err := conn.Where("room_id = ? and player_name = ? ", RoomId, CustName).First(&newPlayer).Error; err != nil {
		log.Println(err.Error())
		return Entity.Result{500, "加入房间失败"}
	}
	if newPlayer != (Entity.Player{}) {
		return Entity.Result{200, "加入房间成功"}
	}
	newPlayer = Entity.Player{RoomId: RoomId, PlayerName: CustName, PlayerRole: Role}
	if err := conn.Save(&newPlayer).Error; err != nil {
		log.Println(err.Error())
		return Entity.Result{500, "加入房间失败"}
	}
	return Entity.Result{200, "加入房间成功"}
}
