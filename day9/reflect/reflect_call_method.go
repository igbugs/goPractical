package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string
}

// 必须使用 public 的函数，才能进行反射否则取不到
func (u *User) Print() {
	fmt.Printf("Print: name: %s age: %d sex: %s\n", u.Name, u.Age, u.Sex)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func TestValue(a interface{}) {
	v := reflect.ValueOf(a)
	//t := reflect.TypeOf(a)

	m := v.MethodByName("Print")
	var args []reflect.Value
	m.Call(args)

	args = args[0:0]	// 清空参数列表
	m = v.MethodByName("SetName")
	args = append(args, reflect.ValueOf("XYB"))
	m.Call(args)
}

func main() {
	var user = User{
		Name: "xyb",
		Age:  12,
		Sex:  "M",
	}

	TestValue(&user)
	fmt.Printf("user: %#v\n", user)
}