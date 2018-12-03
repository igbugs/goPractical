package main

import (
	"fmt"
)

func sendData(ch chan string, e chan bool) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	ch <- "e"
	ch <- "f"
	ch <- "g"
	close(ch)

	e <- true
}

func getData(ch chan string, e chan bool) {
	//var input string
	for {
		input, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(input)
	}
	e <- true
}

func main() {
	ch := make(chan string)
	exitCH := make(chan bool, 2)

	//for i := 0; i < 2; i++ {
	go sendData(ch, exitCH)

	//}

	//for i := 0; i < 2; i++ {
	go getData(ch, exitCH)

	//}

	for i := 0; i < 2; i++ {
		<-exitCH
	}

	//time.Sleep(1 * time.Second)
}
