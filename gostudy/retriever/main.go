package main

import (
	"gostudy/retriever/mock"
	"gostudy/retriever/realcase"
	"fmt"
	"time"
)

const url = "http://www.baidu.com"

type Retriever interface {
	Get(url string) string
	// 定义了Retriever 的接口，接口要存在 Get 的方法 ，参数url 和返回值为string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	// download 是一个接口的使用者(给出我想要的接口的样子)，接口作为参数传入到函数中去
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string {
			"name": "igbugs",
			"course": "goland",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"contents": "another faked baidu.com.",
	})
	return s.Get(url)
}

// type-switch
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *realcase.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println()
}


func main() {
	var r Retriever

	retriever := mock.Retriever{"This is a fake baidu.com"}
	r = &retriever
	inspect(r)

	r = &realcase.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session.")
	fmt.Println(session(&retriever))
}
