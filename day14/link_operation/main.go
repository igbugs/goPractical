package main

import "fmt"

type Person struct {
	Name string
	Age int
	City string
}

func (p *Person) SetName(name string) *Person {
	p.Name = name
	return p
}

func (p *Person) SetAge(age int) *Person {
	p.Age = age
	return p
}

func (p *Person) SetCity(city string) *Person  {
	p.City = city
	return p
}

func (p *Person) Print()  {
	fmt.Printf("Name: %s, Age: %d, City: %s", p.Name, p.Age, p.City)
}

func main()  {
	p := Person{}
	p.SetName("xyb").SetAge(18).SetCity("beijing").Print()
}
