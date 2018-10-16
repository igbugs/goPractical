package main

import "time"

func add() (sum int) {
	for i := 0; i <= 100000; i++ {
		sum += i
	}
	return
}

func main() {
	for i := 0; i <= 10; i++ {
		go add()
	}

	time.Sleep(1 * time.Minute)
}
