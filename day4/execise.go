package main

import "fmt"

func test1()  {
	a := 100
	fmt.Printf("addr of a: %p\n", &a)
}

func test2(p *int) int {
	*p = 200
	return *p
}

func main() {
	test1()

	a := 1000
	fmt.Printf("origin value of a: %d\n", a)
	test2(&a)
	fmt.Printf("changed value of a: %d\n", a)
}
