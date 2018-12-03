package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "C:/GoProject/Go3Project/src/day5/homework/imageThumb/imageThumb.go"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Open file %s failed, err %v\n", filename, err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var content []byte
	var buf [4096]byte

	for {
		n, err := reader.Read(buf[:])
		if err != nil && err != io.EOF {
			fmt.Printf("read file %s failed, err: %v\n", filename, err)
		}

		vaildBuf := buf[0:n]
		content = append(content, vaildBuf...)

		fmt.Printf("content: %s\n", content)

		if err == io.EOF {
			break
		}

	}
}
