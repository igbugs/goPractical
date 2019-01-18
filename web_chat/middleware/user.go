package middleware

import (
	"github.com/gin-gonic/gin"
	"logging"
	"web_chat/common/session"
)

const (
	UserIDkey = "user_id"
	UserNameKey = "username"
)

func UserMiddleware(ctx *gin.Context) {
	sess, err := session.GlobalSess.SessionStart(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.Set(UserIDkey, int64(0))
		return
	}

	var userId int64
	var username string
	var ok bool

	defer func() {
		logging.Debug("user_id:%s username:%s", userId, username)
	}()

	userId, ok = sess.Get(UserIDkey).(int64)
	if !ok {
		ctx.Set(UserIDkey, int64(0))
		return
	}

	ctx.Set(UserIDkey, int64(userId))

	username, ok = sess.Get(UserNameKey).(string)
	if !ok {
		return
	}

	ctx.Set(UserNameKey, username)
	sess.Set(UserIDkey, userId)
	sess.Set(UserNameKey, username)
	defer sess.SessionRelease(ctx.Writer)
	ctx.Next()
}

//func IsLogin(ctx *gin.Context) bool {
//	userId := GetUserId(ctx)
//	if userId <= 0 {
//		return false
//	}
//
//	return true
//}
//
//func GetUserId(ctx *gin.Context) (userId int64) {
//	tmp, exists := ctx.Get(UserIDkey)
//	if !exists {
//		return
//	}
//
//	var ok bool
//	userId, ok = tmp.(int64)
//	if !ok {
//		return
//	}
//
//	return
//}
//
//func GetUserName(ctx *gin.Context) (username int64) {
//	tmp, exists := ctx.Get(UserNameKey)
//	if !exists {
//		return
//	}
//
//	var ok bool
//	username, ok = tmp.(int64)
//	if !ok {
//		return
//	}
//
//	return
//}
