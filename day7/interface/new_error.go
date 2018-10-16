package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	//牛顿法求二次根
	z := float64(1)
	for {
		y := z - (z*z-x)/(2*z)
		if math.Abs(y-z) < 1e-10 {
			return y, nil
		}
		z = y
	}
	return z, nil
}
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
