package api

import (
	"NiuGame/main/Auth"
	"NiuGame/main/Config"
	"NiuGame/main/api/service"
	"NiuGame/main/tool"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	Register(&customer{})
}

func (c *customer) Login(context *gin.Context) {
	//接口接收对象
	var custm customer
	err := context.Bind(&custm)
	//数据库查询对象
	var user customer
	conn.Where(&custm).First(&user)
	//判断结构体对象为空
	if user == (customer{}) {
		tool.ErrHandle(context, "用户名不存在或密码错误，请重试")
		return
	}
	//生成token
	//获取jwt配置
	jwtConfig := Config.GetConfig().JwtConfig
	token, err := Auth.GenerateJwtToken(jwtConfig.SecretKey, jwtConfig.Issuer, jwtConfig.Audience,
		jwtConfig.Expires, custm.Name)
	if err != nil {
		tool.ErrHandle(context, "登录失败")
		return
	}
	//解析token
	claims, err := Auth.ParseJwtToken(token, jwtConfig.SecretKey)
	if err != nil {
		tool.UnAuthHandle(context)
		return
	}
	tool.ResponseHandle(context, "登录成功", claims)
}

func (c *customer) BuildGame(context *gin.Context) {
	var roomRequestDto room
	err := context.Bind(&roomRequestDto)
	if err != nil {
		tool.ErrHandle(context, "创建房间失败")
		log.Panicln(err.Error())
	}
	var roomPo room
	conn.Where("room_id = ? and enable = true", roomRequestDto.RoomId).First(&roomPo)
	if roomPo != (room{}) {
		tool.ErrHandle(context, "创建房间失败，房间号已存在")
		return
	} else {
		//创建房间
		jwtUser, _ := context.Get("jwtUser")
		if jwtUser != nil {
			var claims *Auth.CustomerClaims
			claims = jwtUser.(*Auth.CustomerClaims)
			if claims == nil {
				return
			}
			roomRequestDto.BelongCust = claims.UserName
			conn.Save(&roomRequestDto)
			tool.ResponseHandle(context, "房间创建成功", roomRequestDto.RoomId)
		} else {
			tool.UnAuthHandle(context)
		}
	}
}

func (c *customer) RoomList_get(context *gin.Context) {
	jwtUser, _ := context.Get("jwtUser")
	if jwtUser != nil {
		var claims *Auth.CustomerClaims
		RoomList := []room{}
		claims = jwtUser.(*Auth.CustomerClaims)
		conn.Where("belong_cust = ? and enable = true", claims.UserName).Find(&RoomList)
		tool.ResponseHandle(context, "成功", RoomList)
	} else {
		tool.UnAuthHandle(context)
	}
}

func (c *customer) CloseRoom(context *gin.Context) {
	jwtUser, _ := context.Get("jwtUser")
	if jwtUser != nil {
		var claims *Auth.CustomerClaims
		claims = jwtUser.(*Auth.CustomerClaims)
		var roomRequestDto room
		if err := context.Bind(&roomRequestDto); err != nil {
			tool.ErrHandle(context, "失败")
			log.Panicln(err.Error())
		}
		roomRequestDto.BelongCust = claims.UserName
		var roomPo room
		conn.Where(&roomRequestDto).First(&roomPo)
		if roomPo != (room{}) {
			conn.Model(&roomRequestDto).Update("enable", false)
			tool.OkHandle(context, "成功")
		} else {
			tool.ErrHandle(context, "获取房间信息失败，无法删除")
			log.Panicln("获取房间信息失败，无法删除")
		}
	} else {
		tool.UnAuthHandle(context)
	}
}

func (c *customer) JoinRoom(context *gin.Context) {
	jwtUser, _ := context.Get("jwtUser")
	if jwtUser != nil {
		var claims *Auth.CustomerClaims
		claims = jwtUser.(*Auth.CustomerClaims)
		var roomRequestDto room
		if err := context.Bind(&roomRequestDto); err != nil {
			tool.ErrHandle(context, "失败")
			log.Panicln(err.Error())
		}
		roomRequestDto.Enable = true
		var roomPo room
		conn.Where(&roomRequestDto).First(&roomPo)
		if roomPo != (room{}) {
			Role, err := service.GetRole(roomRequestDto.RoomId, claims.UserName)
			if err != nil {
				tool.ErrHandle(context, "未知错误")
				log.Panicln("未知错误")
			}
			//Role := 2
			if Role == 2 {
				if roomRequestDto.RoomPasswd != roomPo.RoomPasswd {
					tool.ErrHandle(context, "房间密码错误")
					log.Panicln("房间密码错误")
				}
			}
			_, count := service.GetPlayersByRoomId(roomRequestDto.RoomId)
			if count >= 8 {
				tool.ErrHandle(context, "房间"+roomRequestDto.RoomId+"满员")
				log.Panicln("房间" + roomRequestDto.RoomId + "满员")
			}
			result := service.AddPlayer(roomRequestDto.RoomId, claims.UserName, Role)
			tool.ResultHandle(context, result)
		} else {
			tool.ErrHandle(context, "获取房间信息失败，无法加入")
			log.Panicln("获取房间信息失败，无法加入")
		}
	} else {
		tool.UnAuthHandle(context)
	}
}
