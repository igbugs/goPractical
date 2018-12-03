package main

import (
	"fmt"
	"os"
)

func Calc(a, b int) (sum, sub int) {
	sum = a + b // 不需要使用 := 进行复制，因为在函数的定义的时候 已经进行了 sum 与sub 的定义
	sub = a - b
	return
}

func testCalc() {
	sum, sub := Calc(3, 6)
	fmt.Printf("sum: %d, sub: %d\n", sum, sub)
}

func Add(a ...int) int {
	fmt.Printf("func args count %d\n", len(a))

	var sum int
	for i, arg := range a {
		fmt.Printf("arg[%d] = %d\n", i, arg)
		sum += arg
	}
	return sum
}

func testAdd() {
	sum := Add()
	fmt.Printf("sum = %d\n", sum)

	sum = Add(1) // sum 不需要使用 := 赋值
	fmt.Printf("sum = %d\n", sum)

	sum = Add(1, 2)
	fmt.Printf("sum = %d\n", sum)
}

func testDefer() {
	//defer fmt.Printf("Hello World!\n")
	//defer fmt.Printf("你好!\n")

	file, err := os.Open("C:/GoProject/Go3Project/src/day3/map/map.go")
	defer func() {
		if err != nil {
			fmt.Printf("open the file is faild! %s\n", err)
			return
		} else {
			file.Close()
		}
	}()
	//if err != nil {
	//	fmt.Printf("open the file is faild! %s\n", err)
	//	//return
	//}
	//defer file.Close()		// 写在 err 判断之后，因为如果 err 不为 nil 则file 可能为空

	var buf [256]byte
	n, err := file.Read(buf[:])
	if err != nil {
		fmt.Printf("read the file is faild! %s\n", err)
		//return
	}

	fmt.Printf("read %d bytes sunccess! content:\n%s\n", n, string(buf[:]))
	return
}

func main() {
	//testCalc()
	//testAdd()
	testDefer()
}
