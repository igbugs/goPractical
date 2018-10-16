package main

import (
	"reflect"
	"fmt"
	)

func testType(i interface{})  {
	t := reflect.TypeOf(i)
	fmt.Printf("i 的类型为: %v\n", t)

	switch t.Kind() {
	case reflect.Int:
		fmt.Printf("i 的类型为 int\n")
	case reflect.String:
		fmt.Printf("i 的类型为 string\n")
	}
}

func testValue(i interface{})  {
	v := reflect.ValueOf(i)

	switch v.Type().Kind() {
	case reflect.Int:
		v.SetInt(1000)
	case reflect.String:
		v.SetString("helllo")
	case reflect.Ptr:
		// v.Elem() 方法获取的是指针指向的变量(e)，使用e.Type() 取得变量的类型(t),
		// 使用t.kind() 取得反射的类型判断
		switch v.Elem().Type().Kind() {
		case reflect.Int:
			v.Elem().SetInt(1000)
		case reflect.String:
			v.Elem().SetString("helllo")
		}
	}
}

func main() {
	var a int
	testType(a)
	var b string
	testType(b)

	testValue(&a)
	testValue(&b)
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %s\n", b)
}
