package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)
	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "API路由不正确.")
	})

	c := engine.Group("/customer")
	{

		c.POST("/login", GlobalHandler.Login)

	}

	r := engine.Group("/room")
	{
		r.POST("/build", GlobalHandler.BuildRoom)
		r.POST("/close", GlobalHandler.CloseRoom)
		r.POST("/join", GlobalHandler.JoinRoom)
		r.GET("/list", GlobalHandler.RoomList)
		r.GET("/exit", GlobalHandler.ExitRoom)
	}

	p := engine.Group("/player")
	{
		p.GET("/start", GlobalHandler.Start)
	}
	return engine
}
