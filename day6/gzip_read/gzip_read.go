package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "C:/GoProject/Go3Project/src/day6/test.txt.gz"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Open file %s failed, err %v\n", filename, err)
	}

	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Printf("Open gzipfile %s failed, err %v\n", filename, err)
	}

	var content []byte
	var buf [4096]byte
	//r := bufio.NewReader(reader)

	for {
		n, err := reader.Read(buf[:])
		if err != nil && err != io.EOF {
			fmt.Printf("read file %s failed, err: %v\n", filename, err)
		}

		if n > 0 {
			vaildBuf := buf[0:n]
			content = append(content, vaildBuf...)
			fmt.Printf("content: %s\n", content)
		}

		if err == io.EOF {
			break
		}

	}

}
