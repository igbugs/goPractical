package main

import (
	"fmt"
	"time"
)

func main() {
	var intChan chan int = make(chan int)
	fmt.Printf("%p\n", intChan)

	go func() {
		fmt.Printf("insert channel 100\n")
		intChan <- 100
	}()

	go func() {
		fmt.Printf("read from channel\n")
		var a int
		a = <- intChan
		fmt.Printf("a = %d\n", a)
	}()

	time.Sleep(1 * time.Second)
}
