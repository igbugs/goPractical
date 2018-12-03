package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [...]int{1, 3, 5, 8, 7, 2, 3, 4, 6, 4}
	str := [...]string{"A", "ryf", "bbb", "^T@"}
	var si []int
	var ss []string
	si = arr[:]
	ss = str[:]

	sort.Ints(si)
	fmt.Println(si)

	sort.Strings(ss)
	fmt.Println(ss)

}
