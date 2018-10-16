package main

import (
	"fmt"
)


func main() {
	//var c int
	//c = calc.Add(2, 3)
	//
	//fmt.Printf("c = %d\n", c)
	//
	//fmt.Println(math.MaxFloat32)
	//fmt.Println(math.MaxInt8)
	//
	//fmt.Println(math.Pow(2, 32))

	var a, b byte = 6, 11

	//var a1, b1 byte = 'a', 'b'
	//var ai, bi int = 97, 98
	//var aa int = 6
	//var aa1 int32 = 6
	//var aaaa int64 = 6
	//var aaa byte = 6

	//fmt.Println(^a1)
	//fmt.Println(a1 & b1)
	//fmt.Println(^ai)
	//fmt.Println(ai & bi)
	//fmt.Println(^a)
	//fmt.Println(^aa)
	//fmt.Println(^aa1)
	//fmt.Println(^aaaa)
	//fmt.Println(^aaa)
	fmt.Println(a &^ b)
	fmt.Println(b &^ a)

	var c, d int = 15, 6
	var f float64
	f = float64(c /d)
	fmt.Println(f)


}

