package main

import (
	"NiuGame/main/Auth"
	"NiuGame/main/Config"
	"NiuGame/main/api"
	"NiuGame/main/common"
	"NiuGame/main/db"
	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnInit()
}

func main() {

	_, err := Config.ParseConfig(common.File_App_Json)
	if err != nil {
		panic("读取配置文件失败，" + err.Error())
	}
	r := gin.Default()
	r.Use(Auth.JWTAuth())
	// 用户登录接口
	r.POST("/user/login", api.Login)
	r.GET("/getCards", getCards)
	_ = r.Run()
}
