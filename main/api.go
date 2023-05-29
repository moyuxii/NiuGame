package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getCards(c *gin.Context) {
	var userList []User
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
