package main

import "fmt"

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
	// 闭包的几个特征：
	// 1. 返回的是一个匿名函数
	// 2. 这个匿名函数会引用其外部(超出其函数定义域范围)的变量
	// 3. 闭包是相当于将一个匿名函数的外部变量与这个匿名函数 包裹在一起 返回
}
func main() {
	var f = Adder()
	fmt.Printf("Call f(1) = %d\n", f(1))
	fmt.Printf("Call f(22) = %d\n", f(22))
	fmt.Printf("Call f(100) = %d\n", f(100))
}
