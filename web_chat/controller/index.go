package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_chat/common/userinfo"
)

func IndexHandle(ctx *gin.Context) {
	logined := userinfo.IsLogin(ctx)
	if !logined {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		return
	}

	roomList, err := roomMgr.GetRoomList()
	if err != nil {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
		return
	}

	ctx.HTML(http.StatusOK, "views/index.html", roomList)
}
