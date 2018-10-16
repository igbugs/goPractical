package main

import "fmt"

type user struct {
	Name string
	Age int
}

func test1()  {
	var p *int = new(int)
	*p = 1000
	fmt.Printf("*p: %v\naddr: %v\n", *p, p)

	var pUser *user = new(user)
	(*pUser).Name = "xyb"
	(*pUser).Age = 12
	fmt.Printf("*pUser: %#v\n", *pUser)

	pUser.Name = "xyb2"
	pUser.Age = 122
	fmt.Printf("pUser: %#v\n", *pUser)
}

func test2() {
	var p *[]int = new([]int)
	*p = make([]int, 10)

	(*p)[0] = 100
	(*p)[1] = 200
	fmt.Printf("p: %#v\n", *p)

	var m *map[string]int = new(map[string]int)
	*m = make(map[string]int, 10)
	(*m)["key1"] = 100
	(*m)["key2"] = 1000
	fmt.Printf("m: %#v\n", *m)

}

func main() {
	test1()
	test2()
}
