package main

import (
	"fmt"
	"bytes"
)

func f(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func f1(s string, str ...string) {
	fmt.Println(str)
	for _,ss := range str {
		fmt.Println(s, ss)
	}

}

func main() {
	sum, sub := f(3, 6)
	fmt.Println(sum, sub)
	f1("hello", "li", "xue", "cao")

	data := []byte("中华人民共和国")
	rd := bytes.NewReader(data)

	//r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := rd.Read(buf[:])
	//n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)

	//wr := bytes.NewBuffer(nil)
	//w := bufio.NewWriter(wr)
}

