package main

import (
	"fmt"
	"math"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	//fmt.Printf("pc: %v\n", pc)
	for i := range pc {
		//fmt.Printf("i: %v ",i)
		//fmt.Printf("pc[i/2]: %v ",pc[i/2])
		//fmt.Printf("i&1: %v ",i&1)
		//fmt.Printf("byte(i&1): %v ",byte(i&1))
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Printf("pc[i]: %v\n", pc[i])
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	//fmt.Println(byte(x>>(0*8)))
	//fmt.Println(byte(x>>(1*8)))
	//fmt.Println(byte(x>>(2*8)))
	//fmt.Println(byte(x>>(3*8)))
	//fmt.Println(byte(x>>(4*8)))
	//fmt.Println(byte(x>>(5*8)))
	//fmt.Println(byte(x>>(6*8)))
	//fmt.Println(byte(x>>(7*8)))
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main()  {

	var a uint64 = 24
	b := PopCount(a)
	fmt.Println(b)

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
		fmt.Printf("x = %d e^x = %g\n", x, math.Exp(float64(x)))
	}

	//var a []int
	q := make([]int, 3)[3:]
	fmt.Printf("%#v", q)

	m := make(map[rune]string)
	p := &m
	fmt.Printf("%#v", p)

	m['q'] = "wer"
	fmt.Printf("%#v", m)
}
