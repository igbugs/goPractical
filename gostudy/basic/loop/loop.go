package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func memberTest() {
	slice := []int{0, 1, 2, 3}
	for _, member := range slice {
		fmt.Println(member)
	}

}

func main() {
	//fmt.Println(
	//	convertToBin(5),
	//	convertToBin(13),
	//	convertToBin(13345456),
	//)

	printFile("abc.txt")
	//forever()
	//memberTest()

	// 打印多行的字符串
	fmt.Println()
	s := `sadjk
kkk
12edf435
%^&*()
8ujnewdyh`
	printFileContents(strings.NewReader(s))
}
