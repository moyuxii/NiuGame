package tool

import (
	"NiuGame/main/Entity"
	"NiuGame/main/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResultHandle(context *gin.Context, result Entity.Result) {
	context.JSON(
		http.StatusOK,
		gin.H{"Result": result},
	)
}

func ErrHandle(context *gin.Context, err string) {
	context.JSON(
		http.StatusOK,
		gin.H{"Result": Entity.Result{common.ResultError, err}},
	)
}

func OkHandle(context *gin.Context, msg string) {
	context.JSON(
		http.StatusOK,
		gin.H{"Result": Entity.Result{common.ResultOk, msg}},
	)
}

func ResponseHandle(context *gin.Context, msg string, data interface{}) {
	context.JSON(
		http.StatusOK,
		gin.H{"Result": Entity.Result{common.ResultOk, msg}, "Data": data},
	)
}

func UnAuthHandle(context *gin.Context) {
	context.JSON(
		http.StatusUnauthorized,
		gin.H{"Result": Entity.Result{common.ResultError, "用户信息获取失败"}},
	)
}
