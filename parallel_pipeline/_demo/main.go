package main

import (
	"parallel_pipeline/pipeline"
	"fmt"
	"os"
	"bufio"
	)

func main() {
	const (
		filename = "large.in"
		number   = 100000000
	)

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(number)
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func mergeSource() {
	p1 := pipeline.InMemSort(pipeline.ArraySource(2, 45, 5, 7, 22, 20, 20, 77))
	p2 := pipeline.InMemSort(pipeline.ArraySource(2, 33, 43, 21, 100, 1))
	p := pipeline.Merge(p1, p2)
	for v := range p {
		fmt.Println(v)
	}
}
