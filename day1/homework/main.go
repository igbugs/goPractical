package main

import (
	"fmt"
	"reflect"
)

type student struct {
	id int
	name string
}


func main() {
	var s string = "hello GO"
	var i int = 100

	fmt.Printf("%s\n", s)
	fmt.Printf("%d\n", i)
	fmt.Printf("%b\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%f\n", float32(i))

	var stu student = student{1, "li"}

	t := reflect.TypeOf(stu)
	k := reflect.Struct
	fmt.Println(t.Kind())
	fmt.Println(t.NumField())
	fmt.Println(t.Field(1))
	fmt.Println(k)

}
