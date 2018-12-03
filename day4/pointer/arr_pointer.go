package main

import "fmt"

func modify(arr *[3]int) {
	(*arr)[0] = 90
}
func main() {
	a := [3]int{100, 30, 200}
	modify(&a)
	fmt.Println(a)
}
