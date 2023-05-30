package api

import (
	"NiuGame/main/Auth"
	"NiuGame/main/Entity"
	"NiuGame/main/common"
	"NiuGame/main/db"
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
		c.JSON(
			http.StatusOK,
			gin.H{"Result": Entity.Result{common.ResultError, "创建房间失败"}},
		)
		return
	}
	var roomPo room
	dbp := db.GetDb()
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
		if jwtUser != nil {
			var claims *Auth.CustomerClaims
			claims = jwtUser.(*Auth.CustomerClaims)
			if claims == nil {
				return
			}
			roomRequestDto.BelongCust = claims.UserName
			dbp.Save(&roomRequestDto)
			c.JSON(
				http.StatusOK,
				gin.H{"Result": Entity.Result{common.ResultOk, "房间创建成功"},
					"Data": roomRequestDto.RoomId},
			)
		} else {
			c.JSON(
				http.StatusOK,
				gin.H{"Result": Entity.Result{common.ResultError, "用户信息获取失败"}},
			)
			return
		}
	}
}

func (p *player) RoomList_get(c *gin.Context) {
	jwtUser, _ := c.Get("jwtUser")
	if jwtUser != nil {
		var claims *Auth.CustomerClaims
		RoomList := []room{}
		claims = jwtUser.(*Auth.CustomerClaims)
		dbp := db.GetDb()
		dbp.Where("belong_cust = ?", claims.UserName).Find(&RoomList)
		c.JSON(
			http.StatusOK,
			gin.H{"Result": Entity.Result{common.ResultOk, "成功"},
				"Data": RoomList},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"Result": Entity.Result{common.ResultError, "用户信息获取失败"}},
		)
	}
}
