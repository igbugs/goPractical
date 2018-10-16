package main

import (
	"os"
	"fmt"
	"bufio"
)

func isFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func main() {
	filename := "C:/GoProject/Go3Project/src/day6/test.txt"

	var file *os.File
	var err error

	if isFileExists(filename) {
		// mac need os.O_WRONLY
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0755)
	} else {
		file, err = os.Create(filename)
	}

	if err != nil {
		fmt.Printf("Open file %s failed, err %v\n", filename, err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("hohohohoho")

	writer.Flush()
}
