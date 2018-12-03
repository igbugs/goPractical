package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type Result struct {
//	Message string	`json:"message"`
//	Code int		`json:"code"`
//}
//
//type UserInfo struct {
//	Result
//	UserName string `json:"username"`
//	Passwd string 	`json:"passwd"`
//}

func handleFileUpload(c *gin.Context) {
	file, err := c.FormFile("testfile")
	if err != nil {
		fmt.Printf("upload file failed")
		return
	}

	filename := fmt.Sprintf("C:/GoProject/Go3Project/%s", file.Filename)
	err = c.SaveUploadedFile(file, filename)
	if err != nil {
		fmt.Printf("save file failed, err: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, "file upload succ")
}

func handleMultiFileUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Printf("upload file failed")
		return
	}

	files := form.File["testfile"]
	for _, file := range files {
		filename := fmt.Sprintf("C:/GoProject/Go3Project/%s", file.Filename)
		err := c.SaveUploadedFile(file, filename)
		if err != nil {
			fmt.Printf("save file failed, err: %v\n", err)
			return
		}
	}

	c.JSON(http.StatusOK, "file upload succ")
}

func main() {
	r := gin.Default()
	r.POST("/file/upload", handleFileUpload)
	r.POST("/files/upload", handleMultiFileUpload)

	r.Run(":9090")
}
