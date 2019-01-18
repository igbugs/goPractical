package routers

import (
	"github.com/gin-gonic/gin"
	"web_chat/controller"
	"web_chat/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.UserMiddleware)

	r.LoadHTMLGlob("/template/*")
	r.Static("/static/", "/static/")

	r.GET("/index", controller.IndexHandle)
	r.GET("/user/login", loginHandle)
	r.POST("/user/register", registerHandle)
	r.GET("/room/enter", roomEnterHandle)
	r.GET("/ws", wsHandle)

	return r
}