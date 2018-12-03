package main

import (
	"fmt"
	"github.com/sony/sonyflake"
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
