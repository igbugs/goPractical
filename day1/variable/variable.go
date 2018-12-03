package main

import (
	"fmt"
)

var a int
var b string
var c bool
var d int = 8
var e string = "hello"

var (
	aa int
	bb string
	cc bool
	dd int    = 88
	ee string = "hello world"
)

const (
	a1 = 3 + iota
	b1
	c1
	d1
)

var aaaa = 124325

func main() {
	aaa := 123
	bbb := "hahaha"
	ccc := true
	a = 11
	aa = 22
	b = "aaa"
	bb = "eeee"
	c = false
	cc = true
	fmt.Println(a)
	fmt.Println(aa)
	fmt.Println(b)
	fmt.Println(bb)
	fmt.Println(c)
	fmt.Println(cc)
	fmt.Println(e)
	fmt.Println(ee)
	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)

	fmt.Println(aaaa)

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
}
