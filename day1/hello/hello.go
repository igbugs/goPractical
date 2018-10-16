package main

import (
	"fmt"
)

func main() {
	fmt.Print("hello world")
	fmt.Println()
	fmt.Printf("hello world, %d\n", 100)
	fmt.Println("hello world")

	var a int
	var b int
	fmt.Scan(&a)
	fmt.Println(a)
	bb, _ := fmt.Scanf("%d\n", &b)

	fmt.Println(b)
	fmt.Println(bb)
	//time.Sleep(time.Second)

	//Test()
}
