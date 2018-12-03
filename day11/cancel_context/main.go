package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(&wg, ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
}
