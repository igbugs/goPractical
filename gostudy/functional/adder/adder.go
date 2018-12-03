package main

import (
	"fmt"
)

//func adder() func(int) int {
//	sum := 0
//	return func(v int) int {
//		sum += v
//		return sum
//	}
//}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	//a := adder()     // 直接返回函数 func(v int) int ，此函数有参数 v
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	//}

	fmt.Println()
	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var ba int
		ba, a2 = a2(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, ba)
	}
}
