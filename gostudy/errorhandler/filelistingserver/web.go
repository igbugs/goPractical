package main

import (
	"gostudy/errorhandler/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 定义 HandleFileList 函数为appHandler 类型的函数

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	// errWrapper 增加函数错误处理的功能，返回的还是 http.HandleFunc 要求传入的函数类型
	// errWrapper 与 他的返回的函数的区别是，返回的函数没有 error 的返回值
	return func(writer http.ResponseWriter, request *http.Request) {
		// 进行函数的return，错误处理的逻辑，加进来

		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request) // filelisting.HandleFileList  的返回值为 error 类型
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())
			// 进行日志的输出

			if userError, ok := err.(userError); ok {
				// 判断错误的类型，是userError 的话。输出userError.Message() 的信息
				http.Error(writer,
					userError.Message(),
					http.StatusBadRequest) // 400
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound // 404 , Not Found
			case os.IsPermission(err):
				code = http.StatusForbidden // 403
			default:
				code = http.StatusInternalServerError // 500, 不知道的错误
			}
			http.Error(
				writer,                // 向哪个的 response 汇报错误
				http.StatusText(code), // 一个字符串
				code)                  // 报错的状态码
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))
	// 对 filelisting.HandleFileList 函数进行错误处理功能的统一包装

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
