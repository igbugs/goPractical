package main

import "fmt"

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func main() {
	func1 := add(10)
	// 第一次调用func1 更改了base 的变量 由 10 -> 11, 第二次 传入 i= 2, 则输出的返回的（base）为13
	fmt.Println(func1(1), func1(2))
	func2 := add(100)
	fmt.Println(func2(1), func2(2))
}
