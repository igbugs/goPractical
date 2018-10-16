package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Printf("catch panic exception, err%s\n", err)
			}
		}()

		var p *int
		*p = 10
		fmt.Printf("test nil pointer")
	}()


	for {
		fmt.Printf("next running!\n")
		time.Sleep(time.Second)
	}
}
