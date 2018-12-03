package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println(num)

	for i := 0; i <= 8; i++ {
		go func() {
			for {

			}
		}()
	}
	time.Sleep(15 * time.Second)
}
