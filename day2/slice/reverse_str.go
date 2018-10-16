package slice

import "fmt"

func main() {
	var str string = "abcdefgh"
	bytes := []byte(str)

	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
	}

	str = string(bytes)
	fmt.Printf("reverse strings: %s\n", str)
}
