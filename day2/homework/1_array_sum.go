package main

import (
	"fmt"
)

func main() {
	var sum = 0
	arr := [...]int{2, 4, 6, 8, 10}

	//for i := 0; i < len(arr); i++ {
	//	sum += arr[i]
	//}

	for _, v := range arr {
		sum += v
	}

	fmt.Printf("数组的所有元素的和为：%d", sum)
}
