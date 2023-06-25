package handler

import (
	"NiuGame/main/Auth"
	"NiuGame/main/common"
	"NiuGame/main/model"
	"NiuGame/main/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type GlobalHandler struct {
	Src *service.GlobalService
}

/*
 * 登录
 * Param name 账号
 * Param passwd 密码
 * return 用户token
 */
func (c *GlobalHandler) Login(context *gin.Context) {
	name := context.PostForm("name")
	passwd := context.PostForm("passwd")
	if !c.Src.CheckCustomerPasswd(name, passwd) {
		common.ErrHandle(context, "用户名不存在或密码错误，请重试")
	} else {
		jwtConfig := common.JwtConfig{
			Issuer:    viper.GetString("jwt_config.issuer"),
			Audience:  viper.GetString("jwt_config.audience"),
			Expires:   viper.GetInt64("jwt_config.expires"),
			SecretKey: viper.GetString("jwt_config.secret_key"),
		}
		token, _ := Auth.GenerateJwtToken(jwtConfig.SecretKey, jwtConfig.Issuer, jwtConfig.Audience, jwtConfig.Expires, name)
		claims, err := Auth.ParseJwtToken(token, jwtConfig.SecretKey)
		if err != nil {
			common.ErrHandle(context, err.Error())
			return
		}
		common.ResponseHandle(context, "登录成功", map[string]interface{}{"token": token, "info": claims})
	}
}

/*
 * 创建房间
 * Param model.Room
 * return
 */
func (c *GlobalHandler) BuildRoom(context *gin.Context) {
	var roomRequset model.Room
	_ = context.Bind(&roomRequset)
	var err error
	if roomRequset != (model.Room{}) {
		err = c.Src.AddRoom(roomRequset)
		if err != nil {
			common.OkHandle(context, "成功")
		}
	}
	common.ErrHandle(context, err.Error())
}

/*
 * 房间列表
 * Param 空
 * Return []model.Room
 */
func (c *GlobalHandler) RoomList(context *gin.Context) {
	roomList := c.Src.RoomListByAll()
	common.ResponseHandle(context, "成功", roomList)
}

/*
 * 关闭房间
 * Param model.room
 * Return
 */
func (c *GlobalHandler) CloseRoom(context *gin.Context) {
	loginUser, _ := context.Get("LoginUser")
	UserName := loginUser.(string)
	var roomRequestDto model.Room
	_ = context.Bind(&roomRequestDto)
	var err error
	if roomRequestDto != (model.Room{}) {
		err = c.Src.DeleteRoom(roomRequestDto, UserName)
		if err == nil {
			common.OkHandle(context, "成功")
		}
		common.ErrHandle(context, "失败")
	}
}

/*
 * 加入房间
 * Param model.Room
 * Return
 */
func (c *GlobalHandler) JoinRoom(context *gin.Context) {
	loginUser, _ := context.Get("LoginUser")
	UserName := loginUser.(string)
	var roomRequestDto model.Room
	_ = context.Bind(&roomRequestDto)
	roomRequestDto.Enable = true
	err := c.Src.JoinRoom(roomRequestDto, UserName)
	if err != nil {
		common.ErrHandle(context, err.Error())
	}
	common.OkHandle(context, "成功")
}

func (c *GlobalHandler) ExitRoom(context *gin.Context) {
	loginUser, _ := context.Get("LoginUser")
	UserName := loginUser.(string)
	RoomId := context.Query("room_Id")
	err := c.Src.ExitRoom(RoomId, UserName)
	if err != nil {
		common.ErrHandle(context, err.Error())
	}
	common.OkHandle(context, "成功")
}

func (c *GlobalHandler) Start(context *gin.Context) {
	RoomId := context.Query("room_id")
	err := c.Src.Start(RoomId)
	if err != nil {
		common.ErrHandle(context, err.Error())
	}
	common.OkHandle(context, "成功")
}
