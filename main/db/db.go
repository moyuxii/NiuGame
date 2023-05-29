package db

import (
	"NiuGame/main/Config"
	"NiuGame/main/common"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

var dba *gorm.DB = nil

func ConnInit() {
	var err error
	dba, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panicln("connect to db failed,", err)
		return
	}

	err = dba.AutoMigrate(Config.Customer{})
	if err != nil {
		log.Panicln("autoMigrate failed,", err)
		return
	}
	sqlFile, _ := ioutil.ReadFile(common.FIle_Init_Sql)
	_ = dba.Exec(string(sqlFile))
}

func GetDb() *gorm.DB {
	return dba
}
