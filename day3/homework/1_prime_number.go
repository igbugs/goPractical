package main

import (
	"math"
	"fmt"
)

func isPrime(n int) bool {
	var flag bool = true

	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			flag = false
		}
	}
	return flag
}

func main() {
	for i := 1; i < 100; i++ {
		if i < 2 {
			continue
		} else if isPrime(i) {
			fmt.Println(i)
		}
	}
}
