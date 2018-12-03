package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "post", "hahaha")
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	r.GET("/user/info", handleHtml)

	r.Run(":9090")
}
