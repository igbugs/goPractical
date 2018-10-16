package main

import (
	"fmt"
	"time"
	"strings"
)

type TestFunc func(s string) string

func (t TestFunc) doSome(s string) string {
	return fmt.Sprintf("doSomeThing(return string) called: %s\n", t(s))
}

type TestFunc1 func(s string) TestFunc1

func (t TestFunc1) doSome(s string) string {
	return fmt.Sprintf("doSomeThing(return func) called: %s\n", t(s))
}

func add1(r rune) rune { return r +1 }

func main() {
	var f TestFunc = func(s string) string {
		return fmt.Sprintf("TestFunc called: %s\n\n", s)
	}
	a := f("a: call TestFunc .......")
	b := f.doSome("b: call TestFunc.doSome ......")
	fmt.Println(a)
	fmt.Println(b)

	var f1 TestFunc1 = func(s string) TestFunc1 {
		return func(s string) TestFunc1 {
			fmt.Println(s)
			return nil
		}
	}
	f1("a1: call TestFunc .......").doSome("call TestFunc.doSome ......")
	b1 := f1.doSome("b1: call TestFunc.doSome ......")
	//fmt.Println(a1)
	fmt.Println(b1)

	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	fmt.Println(deadline)


	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"

}
