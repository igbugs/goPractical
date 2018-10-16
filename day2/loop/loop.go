package loop

import "fmt"

func main() {
	for i, no := 1, 10; i <= 10 && no <= 19; i, no = i+1, no +1 {
		fmt.Printf("%d * %d = %d\n", i, no, i*no)
	}

	if num := 10; num % 2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
