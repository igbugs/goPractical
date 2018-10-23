package main

import (
	"context"
	"fmt"
)

func add(ctx context.Context, a, b int) int {
	traceid := ctx.Value("trace_id").(string)
	fmt.Printf("[add] trace_id: %v\n", traceid)
	return a + b
}

func cacl(ctx context.Context, a, b int) int {
	traceid := ctx.Value("trace_id").(string)
	fmt.Printf("[cacl] trace_id: %v\n", traceid)
	return add(ctx, a, b)
}


func main()  {
	ctx := context.WithValue(context.Background(), "trace_id", "123456")
	cacl(ctx, 200, 345)
}
