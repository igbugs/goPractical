package slice

import "fmt"

func main() {
	var char string = "我在学习GO"
	chars := []rune(char)

	for i := 0; i < len(chars)/2; i++ {
		chars[i], chars[len(chars)-1-i] = chars[len(chars)-1-i], chars[i]
	}

	char = string(chars)
	fmt.Printf("reverse strings: %s\n", char)

}
