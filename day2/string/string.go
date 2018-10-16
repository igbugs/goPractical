package main

import "fmt"

func main() {
	var str string
	str = "abc汉字"
	var b []byte = []byte(str)
	var chars []rune = []rune(str)

	fmt.Printf("b = %v, len(str)= %d\n", b, len(str))
	fmt.Printf("chars 长度：%d\n", len(chars))

	fmt.Printf("chars 2-4: %v\n", chars[2:4])
	fmt.Printf("b 2-4: %v\n", b[2:4])
}
