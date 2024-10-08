package handler

import (
	"strconv"

	db "webapi/database"
	"webapi/models"
	"webapi/utils"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		r, err := db.GetUser()
		if err != nil {
			utils.FailedResp(ctx, err.Error())
			return
		}
		utils.SuccessResp(ctx, r)
	} else {
		id, err := strconv.Atoi(id)
		if err != nil {
			utils.FailedResp(ctx, err.Error())
			return
		}
		r, err := db.GetUserDetail(int32(id))
		if err != nil {
			utils.FailedResp(ctx, err.Error())
			return
		}
		utils.SuccessResp(ctx, r)
	}

}

func CreateUserHandler(ctx *gin.Context) {
	err := ctx.Request.ParseForm()
	if err != nil {
		utils.FailedResp(ctx, err.Error())
		return
	}

	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	r, err := db.CreateUser(models.User{
		Name:     name,
		Password: password,
	})
	if err != nil {
		utils.FailedResp(ctx, err.Error())
		return
	}
	utils.SuccessResp(ctx, r)
}

func UpdateUserHandler(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		utils.FailedResp(ctx, string("缺少id值"))
		return
	} else {
		id, err := strconv.Atoi(id)
		if err != nil {
			utils.FailedResp(ctx, err.Error())
			return
		}

		var user models.User
		err = ctx.ShouldBindJSON(&user)
		if err != nil {
			utils.FailedResp(ctx, err.Error())
			return
		}

		r, err := db.UpdateUser(int32(id), user)
		if err != nil {
			utils.FailedResp(ctx, err.Error())
			return
		}

		utils.SuccessResp(ctx, r)
	}
}

func DeleteUserHandler(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		utils.FailedResp(ctx, string("缺少id值"))
		return
	} else {
		id, err := strconv.Atoi(id)
		if err != nil {
			utils.FailedResp(ctx, err.Error())
			return
		}
		r, err := db.UpdateUser(int32(id), models.User{Status: 1})
		if err != nil {
			utils.FailedResp(ctx, err.Error())
			return
		}
		utils.SuccessResp(ctx, r)
	}

}
