package array

import "fmt"

func modify(b [3]int) {
	b[0] = 1000
	return
}

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	var a1 [3]int = [3]int{1, 2, 3}

	a2 := [3]int{1, 2, 3}
	a3 := [...]int{1, 2, 3, 4, 5, 6}

	a4 := [3]int{10}
	a5 := [3]int{2:300}

	fmt.Printf("a1 = %v\n", a1)
	fmt.Printf("a2 = %v\n", a2)
	fmt.Printf("a3 = %v\n", a3)
	fmt.Printf("a4 = %v\n", a4)
	fmt.Printf("a5 = %v\n", a5)

	for index, value := range a3 {
		fmt.Printf("a3[%d] = %d\n",index, value)
	}

	var a6 [3][2]int = [3][2]int {
		{1, 2},
		{3, 4},
		{5, 6},
	}

	for index, row := range a6 {
		fmt.Printf("row: %d value: %v\n", index, row)
	}

	for index, row := range a6 {
		for col, value := range row {
			fmt.Printf("a[%d][%d] = %d\n", index, col, value)
		}
	}

	b := a1
	b[0] = 1000
	fmt.Println(a1[0], b[0])

	modify(a1)
	fmt.Println(a1)
}
