package main

import (
	"sort"
	"fmt"
)

func main() {
	a := []int{4, 54, 32, 89, 11, 33, 45}
	sort.Ints(a)

	for i, v := range a {
		fmt.Println(i, v)
	}
}
