package _switch

import "fmt"

func number(a, b int) int {
	num := a * b
	return num
}

func main() {
	//switch finger := 8; finger {
	//case 1:
	//	fmt.Println("Thumb")
	//case 2:
	//	fmt.Println("Index")
	//case 3:
	//	fmt.Println("Middle")
	//case 4:
	//	fmt.Println("Ring")
	//case 5:
	//	fmt.Println("Pinky")
	//default:
	//	fmt.Println("incorrent finger number")
	//}

	//letter := "i"
	//switch letter {
	//case "a","e","i","o","u":
	//	fmt.Println("it's a vowel")
	//default:
	//	fmt.Println("not a vowel")
	//}

	//num := 75
	//switch {
	//case num > 0 && num <= 50:
	//	fmt.Println("num is > 0 and < 50")
	//case num >= 51 && num <= 100:
	//	fmt.Println("num is > 51 and < 100")
	//case num > 101:
	//	fmt.Println("num > 100")
	//
	//}


	switch num := number(15, 5); {
	case num > 0 && num <= 50:
		fmt.Println("num is > 0 and < 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is > 51 and < 100")
	case num > 101:
		fmt.Println("num > 100")

	}
}
