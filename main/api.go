package main

import (
	"NiuGame/main/Auth"
	"NiuGame/main/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func getCards(c *gin.Context) {
	No, _ := strconv.Atoi(c.Query("userNo"))
	Name := c.DefaultQuery("userName", "Non")
	flag := true
	var userTmp User
	if len(userList) > 0 {
		for _, user := range userList {
			if No == user.userId {
				flag = false
				userTmp = user
			}
		}
	}
	if flag {
		userTmp = User{
			userId:   No,
			userName: Name,
			cards: cards{{1, "club"},
				{2, "Diamond"},
				{3, "Heart"},
				{4, "Spade"},
				{5, "Diamond"}},
		}
		userList = append(userList, userTmp)
	}
	fmt.Println(userTmp)
	c.JSON(http.StatusOK, userTmp.cards)
}
