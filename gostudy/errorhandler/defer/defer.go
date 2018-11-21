package main

import (
	"fmt"
	"gostudy/functional/fib"
	"os"
	"bufio"
)

func tryDefer()  {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("error occurred")
	//fmt.Println(4)

	for i := 0; i < 100; i++ {
		//go func() {
		//	defer fmt.Println(i)
		//}()
		defer fmt.Println(i)

		//defer func() {
		//	fmt.Println(i)
		//}()
		if i == 20 {
			panic("printed too many.")
		}
	}
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	// 使用OpenFile来进行文件的打开，os.O_EXCL 这个flag 会在文件存在时报错

	//err = errors.New("This is a custom error!")  // 自定义的error 类型会触发panic 输出
	if err != nil {
		//panic(err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(
				pathError.Op,
				pathError.Path,
				pathError.Err.Error())
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)   // 先写入到buff中，之后通过flush 之后, 写入文件
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
func main()  {
	//tryDefer()
	writeFile("fib.txt")
}