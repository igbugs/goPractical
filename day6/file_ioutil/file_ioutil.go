package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "C:/GoProject/Go3Project/src/day5/homework/imageThumb/imageThumb.go"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file %s failed, err: %v\n", filename, err)
	}
	fmt.Printf("content: %s\n", content)

}
