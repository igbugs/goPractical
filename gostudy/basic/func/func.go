package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) (int, error)  {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}


func apply(op func(int, int,) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()		// 获取函数的指针
	opName := runtime.FuncForPC(p).Name()	// 获取函数的名字
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int)  {
	*a, *b = *b, *a
}

func swap1(a, b int) (int, int) {
	return b, a
}

func main()  {
	if result, err := eval(100, 34, "//"); err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(100, 22)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 4, 5))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	//swap(&a, &b)
	a, b = swap1(a, b)
	fmt.Println(a, b)

}