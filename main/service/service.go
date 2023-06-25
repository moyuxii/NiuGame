package service

import (
	"NiuGame/main/model"
	"NiuGame/main/repository"
	"errors"
	"github.com/spf13/viper"
	"log"
)

type GlobalService struct {
	CustomerRepo *repository.CustomerRepo
	PlayerRepo   *repository.PlayerRepo
	GameRepo     *repository.GameRepo
	RoomRepo     *repository.RoomRepo
}

func (g *GlobalService) CheckCustomerPasswd(name, passwd string) bool {
	customer := model.Customer{Name: name, Passwd: passwd}
	log.Println("正在登录的用户为：", customer)
	return g.CustomerRepo.CheckCustomerPasswd(customer)
}

func (g *GlobalService) AddRoom(roomRequest model.Room) error {
	roomList := g.RoomListByAll()
	for _, room := range roomList {
		if room.RoomId == roomRequest.RoomId {
			return errors.New("房间号已存在，无法创建")
		}
	}
	g.RoomRepo.BuildRoom(roomRequest)
	return nil
}

func (g *GlobalService) RoomListByAll() []model.Room {
	return g.RoomRepo.RoomListByAll()
}

func (g *GlobalService) RoomListByBelongCust(name string) []model.Room {
	return g.RoomRepo.RoomListByBelongCust(name)
}

func (g *GlobalService) DeleteRoom(roomRequest model.Room, userName string) error {
	roomList := g.RoomListByBelongCust(userName)
	for _, room := range roomList {
		if room.RoomId == roomRequest.RoomId {
			g.RoomRepo.Close(room)
			return nil
		}
	}
	return errors.New("删除失败")
}

func (g *GlobalService) JoinRoom(roomRequest model.Room, userName string) error {
	roomList := g.RoomListByAll()
	for _, room := range roomList {
		if room.RoomId == roomRequest.RoomId {
			newPlayer := model.Player{RoomId: room.RoomId, PlayerName: userName, Lock: false, PlayerRole: 2}
			if len(g.PlayerRepo.GetPlayersByRoomId(room.RoomId)) < 8 {
				if room.BelongCust == userName {
					newPlayer.PlayerRole = 1
				}
				if newPlayer.PlayerRole == 1 || room.RoomPasswd == roomRequest.RoomPasswd {
					g.PlayerRepo.AddPlayer(newPlayer)
					return nil
				}
			}
			break
		}
	}
	return errors.New("删除失败")
}

func (g *GlobalService) ExitRoom(roomId, playerName string) error {
	newPlayer := model.Player{RoomId: roomId, PlayerName: playerName}
	player := g.PlayerRepo.GetPlayer(newPlayer)
	if player != (model.Player{}) {
		if !player.Lock {
			g.PlayerRepo.DeletePlayer(player)
			return nil
		}

	}
	return errors.New("失败")
}

func (g *GlobalService) Start(roomId string) error {
	playerList := g.PlayerRepo.GetPlayersByRoomId(roomId)
	var err error
	if len(playerList) > viper.GetInt("room.startMinSize") {
		itemList := make([]model.Gaming, 0)
		for _, player := range playerList {
			itemList = append(itemList, model.Gaming{RoomId: player.RoomId, PlayerName: player.PlayerName})
		}
		err = g.GameRepo.AddPlayerToGame(itemList)
	} else {
		err = errors.New("人数过少")
	}
	return err
}
