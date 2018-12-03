package main

import "fmt"

type Employee interface {
	Calc() float32
}

type Developer struct {
	Name string
	Base float32
}

func (d *Developer) Calc() float32 {
	return d.Base
}

type PM struct {
	Name   string
	Base   float32
	Option float32
}

func (p *PM) Calc() float32 {
	return p.Base + p.Option
}

type YY struct {
	Name   string
	Base   float32
	Option float32
	Ratio  float32
}

func (y *YY) Calc() float32 {
	return y.Base + y.Option*y.Ratio
}

type EmployeeMgr struct {
	employeelist []Employee
}

func (e *EmployeeMgr) Calc() float32 {
	var sum float32
	for _, v := range e.employeelist {
		sum += v.Calc()
	}

	return sum
}

func (e *EmployeeMgr) AddEmployee(d Employee) {
	e.employeelist = append(e.employeelist, d)
}

func main() {
	var e = &EmployeeMgr{}

	dev := &Developer{
		Name: "dev",
		Base: 10000,
	}
	e.AddEmployee(dev)

	pm := &PM{
		Name:   "pm",
		Base:   10000,
		Option: 1110,
	}
	e.AddEmployee(pm)

	yy := &YY{
		Name:   "yy",
		Base:   10000,
		Option: 1110,
		Ratio:  1.2,
	}
	e.AddEmployee(yy)

	sum := e.Calc()
	fmt.Printf("sum: %f\n", sum)
}
