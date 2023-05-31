package api

import (
	"NiuGame/main/Auth"
	"NiuGame/main/Entity"
	"NiuGame/main/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	Register(&player{})
}

func (p *player) ExitRoom(context *gin.Context) {
	jwtUser, _ := context.Get("jwtUser")
	if jwtUser != nil {
		var claims *Auth.CustomerClaims
		RoomId, _ := strconv.Atoi(context.PostForm("RoomId"))
		var roomPO room
		claims = jwtUser.(*Auth.CustomerClaims)
		conn.Where("belong_cust = ? and enable = true", claims.UserName).Find(&RoomList)
		context.JSON(
			http.StatusOK,
			gin.H{"Result": Entity.Result{common.ResultOk, "成功"},
				"Data": RoomList},
		)
	} else {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{"Result": Entity.Result{common.ResultError, "用户信息获取失败"}},
		)
	}
}
