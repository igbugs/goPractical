package main

import (
	"github.com/sony/sonyflake"
	"fmt"
	"time"
)

func main() {
	st := sonyflake.Settings{}
	sk := sonyflake.NewSonyflake(st)
	for {
		fmt.Println(sk.NextID())
		time.Sleep(time.Second)
	}
}
