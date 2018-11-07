package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 11			// 这些变量成为包内的变量, 没有全局变量的称谓
var bb = "hhhh"
// cc := 55			// 函数的外部不可以使用 冒号的定义方式

var (				// var() 集中定义变量
	ab = 12
	bc = "abc"
	cd = true
)

func variableZeroValue()  {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d  %q\n", a, s)
}

func variableInitialValue()  {
	var a ,b  int = 3, 5
	var s  string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	var a, b, c, d = 3, 5, true, "def"
	var s = "abc"
	fmt.Println(a, b, c, d, s)
}

func variableShorter()  {
	a, b, c, d := 3, 5, true, "def"		// 函数的内部可以使用 冒号的简短定义方式
	b = 40
	s := "abc"
	fmt.Println(a, b, c, d, s)
}

func euler()  {
	fmt.Println(cmplx.Exp(1i * math.Pi) +1)
	fmt.Printf("%.3f\n",cmplx.Exp(1i * math.Pi) + 1 )
}

func triangle() {
	var a, b= 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts()  {
	const (
		filename = "a.txt"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(c)
}

func enums()  {
	const (
		cpp = iota
		_
		python
		goland
		javascripts
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascripts, python, goland)
	fmt.Println(b, kb, mb, gb, tb, pb )
}

func main() {
	//fmt.Println("hello world !!")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	//fmt.Println(aa, bb, ab, bc, cd)
	//
	//euler()
	//triangle()
	//consts()
	enums()
}