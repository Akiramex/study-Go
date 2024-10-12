package controllers

import (
	"a-web-sample/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")

	// redis key命名规范设计
	// 单词与单词之间用:分割
	// article:1:likes 文章1的点赞数
	likeKey := "article:" + articleID + ":likes"

	// Incr方法用于递增key所指定的value，如果value不存在会初始化为0再递增1
	if err := global.RedisDb.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successful",
	})
}

func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeKey := "article:" + articleID + ":likes"

	// Get方法从Redis中拿到value
	likes, err := global.RedisDb.Get(likeKey).Result()

	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}
