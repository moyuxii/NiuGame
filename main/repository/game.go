package repository

import (
	"NiuGame/main/model"
	"log"
)

type GameRepo struct {
	DB model.DataBase
}

func (g *GameRepo) AddPlayerToGame(itemList []model.Gaming) error {
	if err := g.DB.SqlLite.Save(&itemList).Error; err != nil {
		log.Panicln(err.Error())
		return err
	}
	return nil
}
