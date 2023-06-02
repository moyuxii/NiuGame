package api

import (
	"NiuGame/main/Auth"
	"NiuGame/main/api/service"
	"NiuGame/main/tool"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	Register(&player{})
}

func (p *player) ExitRoom(context *gin.Context) {
	jwtUser, _ := context.Get("jwtUser")
	if jwtUser != nil {
		var claims *Auth.CustomerClaims
		RoomId := context.PostForm("RoomId")
		claims = jwtUser.(*Auth.CustomerClaims)
		var playerPO player
		if err := conn.Where("room_id = ? and player_name = ?", RoomId, claims.UserName).First(&playerPO); err != nil {
			tool.ErrHandle(context, "退出失败")
			log.Panicln("退出失败")
		}
		if !playerPO.Lock {
			if err := conn.Delete(player{}, "room_id = ? and player_name = ?", RoomId, claims.UserName); err != nil {
				tool.ErrHandle(context, "退出失败")
				log.Panicln("退出失败")
			}
			tool.OkHandle(context, "成功")
		} else {
			tool.ErrHandle(context, "当前用户被锁定，无法退出房间")
		}
	} else {
		tool.UnAuthHandle(context)
	}
}

func (p *player) StartGame(context *gin.Context) {
	jwtUser, _ := context.Get("jwtUser")
	if jwtUser != nil {
		var claims *Auth.CustomerClaims
		RoomId := context.PostForm("RoomId")
		claims = jwtUser.(*Auth.CustomerClaims)
		Role, _ := service.GetRole(RoomId, claims.UserName)
		if Role != 1 {
			tool.ErrHandle(context, "非房主")
			log.Panicln("非房主")
		}
		if err := conn.Model(player{}).Update("lock", true).Error; err != nil {
			tool.ErrHandle(context, "更新房员状态失败")
			log.Panicln("更新房员状态失败")
		}
		tool.OkHandle(context, "开始！")
	} else {
		tool.UnAuthHandle(context)
	}
}
