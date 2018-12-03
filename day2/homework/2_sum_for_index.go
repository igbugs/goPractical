package main

import "fmt"

func main() {
	arr := [...]int{1, 3, 5, 8, 7, 2, 3, 4, 6, 4}

	var n int
	fmt.Printf("请输入想要的两个数组元素的和：")
	fmt.Scanf("%d\n", &n)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == n-arr[i] {
				fmt.Printf("index: (%d, %d) values: (%d, %d)\n", i, j, arr[i], arr[j])
			} else {
				continue
			}
		}
	}
}
