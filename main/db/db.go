package db

import (
	"NiuGame/main/Config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func ConnInit() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panicln("connect to db failed,", err)
	}

	err = db.AutoMigrate(Config.Customer{})
	if err != nil {
		log.Panicln("autoMigrate failed,", err)
	}

}

func getDbC() *gorm.DB {
	return db
}
