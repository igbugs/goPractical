package main

import "fmt"

func main() {
	var b = false
	var (
		c      = true
		d bool = false
	)

	_ = d
	if !b {
		fmt.Println("b is false")
	}

	if !b && c {
		fmt.Println("result is true")
	}

	if b || c {
		fmt.Println("or operation")
	}
}
