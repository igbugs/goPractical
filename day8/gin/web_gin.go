package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type Result struct {
	Message string	`json:"message"`
	Code int		`json:"code"`
}

type UserInfo struct {
	Result
	UserName string `json:"username"`
	Passwd string 	`json:"passwd"`
}

func handleUserInfo(c *gin.Context) {
	username := c.Query("username")
	passwd := c.DefaultQuery("passwd", "defaultpass")

	var result = UserInfo{
		Result: Result{
			Message: "success",
			Code: 0,
		},
		UserName: username,
		Passwd: passwd,
	}

	c.JSON(http.StatusOK, result)
}

func handleUserParam(c *gin.Context) {
	username := c.Param("username")
	passwd := c.Param("passwd")

	u := c.Params
	uu, ok := u.Get("username")
	if ok {
		fmt.Printf(uu)
	}

	var result = UserInfo{
		Result: Result{
			Message: "success",
			Code: 0,
		},
		UserName: username,
		Passwd: passwd,
	}

	c.JSON(http.StatusOK, result)
}

func main() {
	r := gin.Default()
	r.GET("/user/info", handleUserInfo)
	r.GET("/user/info/:username/:passwd", handleUserParam)


	r.Run(":9090")
}
