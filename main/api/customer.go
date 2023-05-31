package api

import (
	"NiuGame/main/Auth"
	"NiuGame/main/Config"
	"NiuGame/main/Entity"
	"NiuGame/main/api/service"
	"NiuGame/main/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

func (c *customer) BuildGame(context *gin.Context) {
	var roomRequestDto room
	err := context.Bind(&roomRequestDto)
	if err != nil {
		context.JSON(
			http.StatusOK,
			gin.H{"Result": Entity.Result{common.ResultError, "创建房间失败"}},
		)
		return
	}
	var roomPo room
	conn.Where("room_id = ? and enable = true", roomRequestDto.RoomId).First(&roomPo)
	if roomPo != (room{}) {
		context.JSON(
			http.StatusOK,
			gin.H{"Result": Entity.Result{common.ResultError, "创建房间失败，房间号已存在"}},
		)
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
			context.JSON(
				http.StatusOK,
				gin.H{"Result": Entity.Result{common.ResultOk, "房间创建成功"},
					"Data": roomRequestDto.RoomId},
			)
		} else {
			context.JSON(
				http.StatusOK,
				gin.H{"Result": Entity.Result{common.ResultError, "用户信息获取失败"}},
			)
			return
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

func (c *customer) CloseRoom(context *gin.Context) {
	jwtUser, _ := context.Get("jwtUser")
	if jwtUser != nil {
		var claims *Auth.CustomerClaims
		claims = jwtUser.(*Auth.CustomerClaims)
		var roomRequestDto room
		if err := context.Bind(&roomRequestDto); err != nil {
			context.JSON(
				http.StatusOK,
				gin.H{"Result": Entity.Result{common.ResultError, "失败"}},
			)
			log.Panicln(err.Error())
		}
		roomRequestDto.BelongCust = claims.UserName
		var roomPo room
		conn.Where(&roomRequestDto).First(&roomPo)
		if roomPo != (room{}) {
			conn.Model(&roomRequestDto).Update("enable", false)
			context.JSON(
				http.StatusOK,
				gin.H{"Result": Entity.Result{common.ResultError, "成功"}},
			)
		} else {
			context.JSON(
				http.StatusOK,
				gin.H{"Result": Entity.Result{common.ResultError, "获取房间信息失败，无法删除"}},
			)
			log.Panicln("获取房间信息失败，无法删除")
		}
	} else {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{"Result": Entity.Result{common.ResultError, "用户信息获取失败"}},
		)
	}
}

func (c *customer) JoinRoom(context *gin.Context) {
	jwtUser, _ := context.Get("jwtUser")
	if jwtUser != nil {
		var claims *Auth.CustomerClaims
		claims = jwtUser.(*Auth.CustomerClaims)
		var roomRequestDto room
		if err := context.Bind(&roomRequestDto); err != nil {
			context.JSON(
				http.StatusOK,
				gin.H{"Result": Entity.Result{common.ResultError, "失败"}},
			)
			log.Panicln(err.Error())
		}
		roomRequestDto.Enable = true
		var roomPo room
		conn.Where(&roomRequestDto).First(&roomPo)
		if roomPo != (room{}) {
			Role := 2
			if roomPo.BelongCust != claims.UserName {
				if roomRequestDto.RoomPasswd != roomPo.RoomPasswd {
					context.JSON(
						http.StatusOK,
						gin.H{"Result": Entity.Result{common.ResultError, "房间密码错误"}},
					)
					log.Panicln("房间密码错误")
				}
			} else {
				Role = 1
			}
			_, count := service.GetPlayersByRoomId(roomRequestDto.RoomId)
			if count > 8 {
				context.JSON(
					http.StatusOK,
					gin.H{"Result": Entity.Result{common.ResultError, "未知错误"}},
				)
				log.Panicln("未知错误")
			}
			if count == 8 {
				context.JSON(
					http.StatusOK,
					gin.H{"Result": Entity.Result{common.ResultError, "房间" + roomRequestDto.RoomId + "满员"}},
				)
				log.Panicln("房间" + roomRequestDto.RoomId + "满员")
			}
			result := service.AddPlayer(roomRequestDto.RoomId, claims.UserName, Role)
			context.JSON(
				http.StatusOK,
				gin.H{"Result": result},
			)
		} else {
			context.JSON(
				http.StatusOK,
				gin.H{"Result": Entity.Result{common.ResultError, "获取房间信息失败，无法加入"}},
			)
			log.Panicln("获取房间信息失败，无法加入")
		}
	} else {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{"Result": Entity.Result{common.ResultError, "用户信息获取失败"}},
		)
	}
}
