package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
	"gostudy/functional/fib"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {		// 为 intGen 类型的函数 实现Read 方法
	next := g() 						// next 接收 g() 函数的输出值
	if next > 20000 {
		return 0, io.EOF
	}

	s := fmt.Sprintf("%d\n", next)   // 转换 next 为 string 类型
	return strings.NewReader(s).Read(p)
	// func NewReader(s string) *Reader { return &Reader{s, 0, -1} }
	// NewReader 传入 字符串，返回一个Reader(结构体，他实现了Read()的方法)
}

func printFileContents(reader io.Reader) {    // 实现了 Read 方法的函数 intGen 可以转入 供printFileContents 打印
	scanner := bufio.NewScanner(reader)
	// func NewScanner(r io.Reader) *Scanner
	// Reader 接口的实现者 r, 传入，返回一个 Scanner 的结构体，Scanner 有方法Text 读取内容

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()

	printFileContents(f)
}
