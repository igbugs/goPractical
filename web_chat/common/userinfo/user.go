package userinfo

import (
	"github.com/gin-gonic/gin"
	"web_chat/middleware"
)

func IsLogin(ctx *gin.Context) bool {
	userId := GetUserId(ctx)
	if userId <= 0 {
		return false
	}

	return true
}

func GetUserId(ctx *gin.Context) (userId int64) {
	tmp, exists := ctx.Get(middleware.UserIDkey)
	if !exists {
		return
	}

	var ok bool
	userId, ok = tmp.(int64)
	if !ok {
		return
	}

	return
}

func GetUserName(ctx *gin.Context) (username int64) {
	tmp, exists := ctx.Get(middleware.UserNameKey)
	if !exists {
		return
	}

	var ok bool
	username, ok = tmp.(int64)
	if !ok {
		return
	}

	return
}
