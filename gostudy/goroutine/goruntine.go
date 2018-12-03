package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func() {
			for {
				//fmt.Printf("Hello from goroutine %d\n", i)
				a[i]++
				// 每个数组索引的内存块的值进行 加 1，没有机会进行协程之间的切换，如果这个协程不主动的交出控制权，
				// 会一直卡住，始终在那一个的协程里面，占用系统的资源
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
