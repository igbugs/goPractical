package main

import "fmt"

func calc(op func(args ...int)int, op_args ...int) int {
	result := op(op_args...)	// 传参的时候可以将op_args 这个slice进行展开
	fmt.Printf("result = %d\n", result)
	return result
}
func main() {
	calc(func(args ...int) int {
		var sum int
		for i := 0; i < len(args);i++ {
			sum += args[i]
		}
		return sum
	}, 1, 2, 3, 4, 5)

	calc(func(args ...int) int {
		var sum int = 1
		for i := 0; i < len(args);i++ {
			sum *= args[i]
		}
		return sum
	}, 1, 2, 3, 4, 5)
}
