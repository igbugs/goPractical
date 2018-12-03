package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var c byte
	var str string

	//var arr [3][3][3] int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	for i := 0; ; i++ {
		c, _ = r.ReadByte()
		if c == 10 {
			break
		} else {
			w.WriteByte(c)
			w.Flush()

			str += string(c)
		}
	}

	fmt.Println()
	fmt.Println(str)
}
