package db

import (
	"NiuGame/main/Entity"
	"NiuGame/main/common"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var dba *gorm.DB = nil

func init() {
	var err error
	dba, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Panicln("connect to db failed,", err)
	}
	//数据库表初始话
	TableInit()
	//数据初始化
	sqlFile, err := os.ReadFile(common.FIle_Init_Sql)
	_ = dba.Exec(string(sqlFile))
}

func TableInit() {
	if dba == nil {
		log.Panicln("orm entity is nil,db init failed")
	}
	var err error
	//用户表、房间表
	err = dba.AutoMigrate(&Entity.Customer{}, &Entity.Room{}, &Entity.Player{})
	if err != nil {
		log.Panicln("init occur failed", err)
	}
}

func GetDb() *gorm.DB {
	return dba
}
