package main

import "fmt"

type Animal interface {
	Eat()
	Talk()
	Run()
}

type Dog struct {
	name string
}

func (d *Dog) Eat() {
	fmt.Printf("%s is eating.\n", d.name)
}

func (d *Dog) Talk() {
	fmt.Printf("%s is talking.\n", d.name)
}

func (d *Dog) Run() {
	fmt.Printf("%s is running.\n", d.name)
}

type Pig struct {
	name string
}

func (p *Pig) Eat() {
	fmt.Printf("%s is eating.\n", p.name)
}

func (p *Pig) Talk() {
	fmt.Printf("%s is talking.\n", p.name)
}

func (p *Pig) Run() {
	fmt.Printf("%s is running.\n", p.name)
}

func main() {
	var a Animal
	var dog = &Dog{
		name: "旺财",
	}

	a = dog
	a.Eat()
	a.Talk()
	a.Run()

	var pig = &Pig{
		name: "佩奇",
	}

	a = pig
	a.Eat()
	a.Talk()
	a.Run()
}
