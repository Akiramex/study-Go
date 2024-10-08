package router

import (
	"webapi/handler"

	"github.com/gin-gonic/gin"
)

func BuildRoute(g *gin.Engine) *gin.Engine {
	route := g.Group("/training/mobile/api/user")
	{
		route.GET("", handler.GetUserHandler)
		route.PUT("", handler.UpdateUserHandler)
		route.POST("", handler.CreateUserHandler)
		route.DELETE("", handler.DeleteUserHandler)
	}

	return g
}
