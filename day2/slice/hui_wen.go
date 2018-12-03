package slice

import "fmt"

func main() {
	var str string = "abcdcba"
	bytes := []byte(str)

	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
	}

	str1 := string(bytes)
	fmt.Printf("reverse strings: %s\n", str1)

	fmt.Println(str == str1)
	if str == str1 {
		fmt.Printf("字符串%s 是一个回文字符串", str)
	} else {
		fmt.Printf("字符串%s 不是一个回文字符串", str)
	}

	a := "abc"
	b := "abc"

	fmt.Println()
	fmt.Println(a == b)
}
