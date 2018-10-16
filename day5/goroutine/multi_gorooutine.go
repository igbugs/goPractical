package main

import (
	"time"
	"fmt"
)

func number() {
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("%d ", i)
	}
}

func char() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go number()
	go char()

	time.Sleep(3 * time.Second)
}
