package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) SetName(name string) {
	a.Name = name
}

func (a *Animal) SetAge(age int) {
	a.Age = age
}

func (a *Animal) Print() {
	fmt.Printf("a.Name=%s a.age=%d\n", a.Name, a.Age)
}

type Birds struct {
	*Animal // 定义之后实例化的时候，要进行初始化
}

func (b *Birds) Fly() {
	fmt.Printf("%s is Flying\n", b.Name)
}

func main() {
	var b *Birds = &Birds{
		&Animal{}, // 定义之后实例化的时候，要进行初始化,要不然会对空指针进行操作
	}

	b.SetName("birds")
	b.SetAge(11)
	b.Fly()
	b.Print()

	data, err := json.Marshal(b)
	fmt.Printf("marshal result: %s err:%v\n", string(data), err)

	var c Birds
	err = json.Unmarshal(data, &c)
	fmt.Printf("c: %#v\n, err: %#v\n", *c.Animal, err)
}
