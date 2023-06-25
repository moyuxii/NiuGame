package repository

import (
	"NiuGame/main/model"
	"log"
)

type PlayerRepo struct {
	DB model.DataBase
}

func (p *PlayerRepo) GetPlayersByRoomId(roomId string) []model.Player {
	var playerList []model.Player
	p.DB.SqlLite.Where("room_id = ?", roomId).Find(&playerList)
	return playerList
}

func (p *PlayerRepo) GetPlayer(playerReq model.Player) model.Player {
	var player model.Player
	p.DB.SqlLite.Where(&playerReq).First(&player)
	return player
}

func (p *PlayerRepo) AddPlayer(player model.Player) {
	p.DB.SqlLite.Save(player)
}

func (p *PlayerRepo) DeletePlayer(playerReq model.Player) bool {
	if err := p.DB.SqlLite.Delete(model.Player{}, playerReq).Error; err != nil {
		log.Panicln(err)
		return false
	}
	return true
}

func (p *PlayerRepo) UpdateLock(roomId string, lock bool) bool {
	if err := p.DB.SqlLite.Where("room_id = ?", roomId).Update("lock", lock); err != nil {
		log.Panicln(err)
		return false
	}
	return true
}
