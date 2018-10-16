package main

import (
	"fmt"
	"sync"
)

func sendData(ch chan string, wg *sync.WaitGroup) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	ch <- "e"
	ch <- "f"
	ch <- "g"
	close(ch)

	wg.Done()
}

func getData(ch chan string, wg *sync.WaitGroup) {
	for {
		input, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(input)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(2)
	go sendData(ch, &wg)
	go getData(ch, &wg)

	wg.Wait()

}
