package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) { // 判断是否已 suffix 结尾，不是则 把suffix 加上
			return name + suffix
		}
		return name
	}
}

func main() {
	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test"))
	fmt.Println(func1("test1.jpg"))
	fmt.Println(func2("test"))
}
