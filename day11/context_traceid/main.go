package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, ctx context.Context) {
	tracid, ok := ctx.Value("Trace_ID").(string)
	if ok {
		fmt.Printf("trace_id: %v\n", tracid)
	}

LOOP:
	for {
		fmt.Printf("worker: trace_id: %v\n", tracid)
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}

	fmt.Println("worker done")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	//ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	ctx = context.WithValue(ctx, "Trace_ID", "12431526")

	wg.Add(1)
	go worker(&wg, ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
