package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

// 必须使用 public 的函数，才能进行反射否则取不到
func (u *User) Test() {
	fmt.Println("This is a test func")
}

func main() {
	var user = User{
		Name: "xyb",
		Age:  12,
		Sex:  "M",
	}

	v := reflect.ValueOf(&user)
	t := v.Type()
	v.Elem().Field(0).SetString("XYB")
	fmt.Println("user: ", user)
	fmt.Println("method num: ", v.NumMethod())

	for i := 0; i < v.NumMethod(); i++ {
		f := t.Method(i)
		fmt.Printf("%d method, name:%v, type:%v\n", i, f.Name, f.Type)
	}
}
