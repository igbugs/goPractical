package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo1 struct {
	UserName string `form:"username" json:"username"`
	Passwd string `form:"passwd" json:"passwd"`
	Age int `form:"age" json:"age"`
	Sex string `form:"sex" json:"sex"`
}

func handleUserInfo1(c *gin.Context) {
	var userinfo UserInfo1
	err := c.ShouldBind(&userinfo)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, userinfo)
}

func handleUserInfoJson(c *gin.Context) {
	var userinfo UserInfo1
	err := c.ShouldBindJSON(&userinfo)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, userinfo)
}

func handleUserInfoQuery(c *gin.Context) {
	var userinfo UserInfo1
	err := c.ShouldBind(&userinfo)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, userinfo)
}

func main() {
	r := gin.Default()

	v1Group := r.Group("/v1")
	v1Group.POST("/user/info", handleUserInfo1)
	v1Group.POST("/user/infojson", handleUserInfoJson)
	v1Group.GET("/user/info", handleUserInfoQuery)

	r.Run(":9090")
}
