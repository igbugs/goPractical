package _switch

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	var number int

	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(100)
	fmt.Printf("请输入一个数字在 [0-100]之间\n")
	for {
		var input int
		fmt.Scanf("%d\n", &input)
		var flag bool = false

		switch {
		case number > input:
			fmt.Printf("输入的数字太小\n")
		case number == input:
			fmt.Printf("恭喜你，答对了\n")
			flag = true
		case number < input:
			fmt.Printf("输入的数字太大\n")
		}

		if flag {
			break
		}
	}
}
