package main

import "fmt"

type T1 struct {
	T2
	T3
}

type T2 struct {
	T4
	foo int
}

type T3 struct {
	T4
}

type T4 struct {
	foo int
}

func main() {
	t2 := T2{ T4{ 9000 }, 2}
	t3 := T3{ T4{ 3 } }
	fmt.Printf("foo = %d\n", t2.foo)
	fmt.Printf("foo = %d\n", t2.T4.foo)
	fmt.Printf("foo = %d\n", t3.foo)

	t1 := T1{
		t2,
		t3,
	}
	fmt.Printf("foo = %d\n", t1.foo)
	fmt.Printf("foo = %d\n", t1.T2.foo)
	fmt.Printf("foo = %d\n", t1.T2.T4.foo)
	fmt.Printf("foo = %d\n", t1.T3.foo)
}
