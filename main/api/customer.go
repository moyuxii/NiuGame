package api

import (
	"NiuGame/main/Auth"
	"NiuGame/main/Config"
	"NiuGame/main/Entity"
	"NiuGame/main/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type customer Entity.Customer

func init() {
	Register(&customer{})
}

func (c *customer) Login(context *gin.Context) {
	//接口接收对象
	var custm customer
	err := context.Bind(&custm)
	//数据库查询对象
	var user customer
	dbc := db.GetDb()
	dbc.Where(&custm).First(&user)
	//判断结构体对象为空
	if user == (customer{}) {
		context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "用户名不存在或密码错误，请重试"})
		return
	}
	//生成token
	//获取jwt配置
	jwtConfig := Config.GetConfig().JwtConfig
	token, err := Auth.GenerateJwtToken(jwtConfig.SecretKey, jwtConfig.Issuer, jwtConfig.Audience,
		jwtConfig.Expires, custm.Name)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "登录失败"})
		return
	}
	//解析token
	claims, err := Auth.ParseJwtToken(token, jwtConfig.SecretKey)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "token解析失败"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "token": token, "claims": claims})
}
