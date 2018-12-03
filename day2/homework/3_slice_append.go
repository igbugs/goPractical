package main

import "fmt"

func main() {
	var sa = make([]string, 5, 10)

	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
	}

	//for i := 0; i < 5; i++ {
	//	fmt.Println(sa[i])
	//}
	fmt.Println(sa, len(sa)) // [     0 1 2 3 4 5 6 7 8 9], 含有 5个的空格
}
