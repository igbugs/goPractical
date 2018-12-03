package string

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes 我在学习GO语言！" // UTF-8
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println()
	for i, ch := range s { // ch is rune
		fmt.Printf("(%d %X)", i, ch)
	}

	fmt.Println()
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:] // 取第一个字符的长度之后的字符
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}

}
