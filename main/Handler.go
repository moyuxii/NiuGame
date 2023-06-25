package main

import (
	"NiuGame/main/Config"
	"NiuGame/main/handler"
	"NiuGame/main/model"
	"NiuGame/main/repository"
	"NiuGame/main/service"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitViper() {
	if err := Config.Init(""); err != nil {
		panic(err)
	}
}

func InitDB() {
	DB = model.GetSqliteDB()
}

func InitHandler() {
	GlobalHandler = handler.GlobalHandler{
		Src: &service.GlobalService{
			CustomerRepo: &repository.CustomerRepo{
				DB: model.DataBase{
					SqlLite: DB,
				},
			},
			RoomRepo: &repository.RoomRepo{
				DB: model.DataBase{
					SqlLite: DB,
				},
			},
			PlayerRepo: &repository.PlayerRepo{
				DB: model.DataBase{
					SqlLite: DB,
				},
			},
			GameRepo: &repository.GameRepo{
				DB: model.DataBase{
					SqlLite: DB,
				},
			},
		},
	}
}
