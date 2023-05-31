package service

import (
	"NiuGame/main/db"
	"gorm.io/gorm"
)

var conn *gorm.DB

func init() {
	conn = db.GetDb()
}
