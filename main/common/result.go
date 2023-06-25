package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func ErrHandle(context *gin.Context, err string) {
	context.JSON(
		http.StatusOK,
		gin.H{"result": Result{ResultError, err}},
	)
}

func OkHandle(context *gin.Context, msg string) {
	context.JSON(
		http.StatusOK,
		gin.H{"result": Result{ResultOk, msg}},
	)
}

func ResponseHandle(context *gin.Context, msg string, data interface{}) {
	context.JSON(
		http.StatusOK,
		gin.H{
			"result": Result{ResultOk, msg},
			"data":   data,
		},
	)
}

func UnAuthHandle(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{"result": Result{Unauthorized, "用户信息获取失败"}},
	)
}
