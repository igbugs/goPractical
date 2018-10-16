package main

import (
	"fmt"
	"bytes"
)

func test_str1()  {
	var a string = "hello \n \n \n"
	var b = "hello"
	fmt.Printf("a = %v, b = %s\n", a, b)
}

func test_str2()  {
	var a string = `
	ksdfjklf
kfjksdj
ksdj`
	fmt.Printf("b = %s\n", a)
}

func test_char()  {
	var c byte
	var d rune

	c = 'w'
	d = '我'
	fmt.Printf("c = %c\n", c)
	fmt.Printf("d = %c\n", d)
}

func main() {
	test_str1()
	test_str2()
	test_char()


	ss := []byte("   hello,    world!  ")
	fmt.Printf("%q", string(bytes.TrimSpace(ss)))
}

