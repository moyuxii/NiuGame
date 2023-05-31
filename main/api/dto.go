package api

import (
	"NiuGame/main/Entity"
	"NiuGame/main/db"
	"gorm.io/gorm"
)

type customer Entity.Customer
type room Entity.Room
type player Entity.Player

var conn *gorm.DB

func init() {
	conn = db.GetDb()
}
