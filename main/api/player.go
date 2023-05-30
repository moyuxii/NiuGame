package api

import (
	"NiuGame/main/Entity"
	"NiuGame/main/common"
	"NiuGame/main/db"
	"NiuGame/main/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

type player Entity.Player
type room Entity.Room

func init() {
	Register(&player{})
}

func (p *player) BuildGame(c *gin.Context) {
	var roomRequestDto room
	err := c.Bind(&roomRequestDto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "创建房间失败"})
		return
	}
	dbp := db.GetDb()
	var roomPo room
	dbp.Where("room_id = ?", roomRequestDto.RoomId).First(&roomPo)
	if roomPo != (room{}) {
		c.JSON(
			http.StatusOK,
			gin.H{"Result": Entity.Result{common.ResultError, "创建房间失败，房间号已存在"}},
		)
		return
	} else {
		//创建房间
		jwtUser, _ := c.Get("jwtUser")
		claims, err2 := tool.ToClaims(jwtUser)
		if err2 != nil {
			c.JSON(
				http.StatusOK,
				gin.H{"Result": Entity.Result{common.ResultError, err2.Error()},
					"Data": roomRequestDto.RoomId},
			)
			return
		}
		roomRequestDto.BolongCust = claims.UserName
		dbp.Save(&roomRequestDto)
	}
	c.JSON(
		http.StatusOK,
		gin.H{"Result": Entity.Result{common.ResultOk, "房间创建成功"},
			"Data": roomRequestDto.RoomId},
	)
}
