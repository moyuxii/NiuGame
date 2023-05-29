package api

import (
	"NiuGame/main/Auth"
	"NiuGame/main/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	jwtConfig := Config.GetConfig().JwtConfig
	var customer Config.Customer
	err := context.Bind(&customer)
	fmt.Println(customer)
	if customer != Config.GetConfig().Customer {
		context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "用户名/密码错误，请重试"})
		return
	}
	//生成token
	token, err := Auth.GenerateJwtToken(jwtConfig.SecretKey, jwtConfig.Issuer, jwtConfig.Audience,
		jwtConfig.Expires, customer.Name)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "登录失败"})
	}
	//解析token
	claims, err := Auth.ParseJwtToken(token, jwtConfig.SecretKey)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "token解析失败"})
	}
	context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "token": token, "claims": claims})
}
