package main

import "fmt"

func testPointer1() {
	var a *int
	var b int = 200

	fmt.Printf("value of a: %v\n", a)
	a = &b
	fmt.Printf("value of a: %v\n", a)
	fmt.Printf("*a: %v\n", *a)

}

func testPointer2() {
	var a *int
	var b int = 200

	a = &b
	fmt.Printf("value of a: %v\n", a)

	fmt.Printf("addr of b: %v\n", &b)
	fmt.Printf("adde of a: %v\n", &a)

	var c int = 300
	a = &c
	fmt.Printf("*a: %v\n", *a)

}

func testPointer3() {
	var a *int
	var b int = 200

	a = &b
	fmt.Printf("value of a: %v\n", a)

	fmt.Printf("addr of b: %v\n", &b)
	fmt.Printf("adde of a: %v\n", &a)

	*a = 300
	fmt.Printf("b: %v\n", b)

}

func main() {
	testPointer1()
	testPointer2()
	testPointer3()
}
