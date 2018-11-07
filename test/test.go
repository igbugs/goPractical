package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main()  {
	//a := make([]int, 0)
	////a := []int{}
	//for i := 0; i <= 100; i++ {
	//	a = append(a, i)
	//	fmt.Println(len(a), cap(a))
	//}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<10; i++ {
		fmt.Println(r.Int())
	}

}
