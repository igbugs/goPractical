package main

import "fmt"

type integer int

func (i integer) print() {
	fmt.Printf("i: %#v\n", i)
}

func (i *integer) set(b int64) {
	*i = integer(b)
}

func main() {
	var a integer = 10000
	fmt.Printf("a = %#v\n", a)

	var b int64 = 400
	a = integer(b)
	fmt.Printf("a = %#v\n", a)

	a.print()
	a.set(20000)
	a.print()
}
