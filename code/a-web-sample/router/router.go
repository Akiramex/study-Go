package router

import (
	"a-web-sample/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	return r
}
