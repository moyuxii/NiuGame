package main

import (
	"NiuGame/main/Auth"
	"NiuGame/main/handler"
	"flag"
	"github.com/gin-gonic/gin"
)

var (
	GlobalHandler handler.GlobalHandler
)

func init() {
	InitViper()
	InitDB()
	InitHandler()
}

func main() {
	flag.Parse()
	r := gin.New()
	_ = Load(
		r,
		Auth.JWTAuth(),
	).Run()
}
