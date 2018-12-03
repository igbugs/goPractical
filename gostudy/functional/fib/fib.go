package fib

func Fibonacci() func() int { // 定义 fibnacci() 函数的返回值为 intGen 类型的函数
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
