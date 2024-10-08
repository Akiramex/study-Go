package utils

import (
	"net/http"
	"webapi/models"

	"github.com/gin-gonic/gin"
)

func SuccessResp(ctx *gin.Context, data ...interface{}) {
	var respData interface{}
	if len(data) > 0 {
		respData = data[0] // 如果有参数，取第一个参数作为Data
	}
	// 响应结构体
	response := models.Response{
		Code: 0,
		Msg:  "",
		Data: respData,
	}
	// 发送JSON响应
	ctx.JSON(http.StatusOK, &response)
}

func FailedResp(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, &models.Response{
		Code: -1,
		Msg:  msg,
		Data: nil,
	})
}
