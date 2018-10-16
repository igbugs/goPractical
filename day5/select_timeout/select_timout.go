package main

import (
	"fmt"
	"time"
)

func queryDb(ch chan int) {
	time.Sleep(time.Second * 2)		// 执行的时间是2s
	ch <- 100
}

func main() {
	ch := make(chan int)
	t := time.NewTicker(time.Second)	// 定义 1s 超时
	//t := time.NewTimer(time.Second)	// 定义 1s 超时

	go queryDb(ch)

	select {
	case v := <-ch:
		fmt.Println("result", v)
	case <-t.C:
		fmt.Println("timeout")
	}
}
