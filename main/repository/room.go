package repository

import (
	"NiuGame/main/model"
	"log"
)

type RoomRepo struct {
	DB model.DataBase
}

func (r *RoomRepo) RoomListByAll() []model.Room {
	var roomList []model.Room
	r.DB.SqlLite.Where("enable = true").Find(&roomList)
	return roomList
}

func (r *RoomRepo) RoomListByBelongCust(name string) []model.Room {
	var roomList []model.Room
	r.DB.SqlLite.Where("enable = true and belong_cust = ?", name).Find(&roomList)
	return roomList
}

func (r *RoomRepo) BuildRoom(room model.Room) {
	if err := r.DB.SqlLite.Save(&room).Error; err != nil {
		log.Panicln(err)
	}
}

func (r *RoomRepo) Close(room model.Room) {
	if err := r.DB.SqlLite.Where(&room).Update("enable", false).Error; err != nil {
		log.Panicln(err)
	}
}
