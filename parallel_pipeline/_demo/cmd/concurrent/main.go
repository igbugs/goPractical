package main

import (
	"fmt"
	)

func printHello(i int, ch chan string)  {
	for {
		ch <- fmt.Sprintf("hello world! from %d\n", i)
	}

}

func main()  {
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		go printHello(i, ch)
	}

	for {
		msg := <- ch
		fmt.Print(msg)
	}
}
