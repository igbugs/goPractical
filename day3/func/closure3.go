package main

import "fmt"

func calca(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	f1, f2 := calca(10)
	// f1 与 f2 引用的是同一个的外部的变量 base, 随着
	// f1(1),f2(2),f1(3)的调用后的base值（来回更改），为11,9,12
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
	fmt.Println(f1(7), f2(8))
}
