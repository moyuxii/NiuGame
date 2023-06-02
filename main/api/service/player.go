package service

import (
	"NiuGame/main/Entity"
	"NiuGame/main/common"
	"log"
)

func AddPlayer(RoomId string, CustName string, Role int) Entity.Result {
	var newPlayer Entity.Player
	if err := conn.Where("room_id = ? and player_name = ? ", RoomId, CustName).First(&newPlayer).Error; err != nil {
		log.Println(err.Error())
		return Entity.Result{common.ResultError, "加入房间失败"}
	}
	if newPlayer != (Entity.Player{}) {
		return Entity.Result{common.ResultOk, "加入房间成功"}
	}
	newPlayer = Entity.Player{RoomId: RoomId, PlayerName: CustName, PlayerRole: Role}
	if err := conn.Save(&newPlayer).Error; err != nil {
		log.Println(err.Error())
		return Entity.Result{common.ResultError, "加入房间失败"}
	}
	return Entity.Result{common.ResultOk, "加入房间成功"}
}

func GetRole(RoomId string, CustName string) (int, error) {
	var count int64
	err := conn.Where("room_id = ? and belong_cust = ?", RoomId, CustName).First(Entity.Room{}).Count(&count).Error
	switch count {
	case 0:
		return 2, err
	case 1:
		return 1, err
	default:
		return 2, err
	}
}
