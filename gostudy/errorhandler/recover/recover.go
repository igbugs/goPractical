package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(fmt.Sprintf("It's not a error type: ", r))
		}
	}()
	//panic(errors.New("This is an error."))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	panic(123) // 直接panic(123), 然后让recover进行处理，recover 会在进行panic
}

func main() {
	tryRecover()
}
