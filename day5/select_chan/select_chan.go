package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var intCh = make(chan int, 10)
	var strCh = make(chan string, 10)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		var count int
		for count < 10000 {
			count++
			select {
			case intCh <- 10:
				fmt.Printf("write to int channel success.\n")
			case strCh <- "a":
				fmt.Printf("write str channel success.\n")
			default:
				fmt.Printf("all channel is full\n")
				time.Sleep(time.Second)
			}
		}
		wg.Done()
	}()

	go func() {
		var count int
		for count < 10000 {
			count++
			select {
			case a := <-intCh:
				fmt.Printf("read from int channel %d.\n", a)
			case a := <-strCh:
				fmt.Printf("read from str channel %s.\n", a)
			default:
				fmt.Printf("all channel is read done\n")
				time.Sleep(time.Second)
			}
		}
		wg.Done()
	}()

	wg.Wait()
}
