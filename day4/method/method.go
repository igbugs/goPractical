package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func (s Student) GetName() string {
	return s.Name
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = Student{
		Name: "s1",
		Age: 100,
	}

	name := s1.GetName()
	fmt.Printf("name = %v\n", name)

	//(&s1).SetName("s2")
	s1.SetName("s2")
	name = s1.GetName()
	fmt.Printf("name = %v\n", name)
}
