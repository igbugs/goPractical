package main

import (
	"fmt"
	"sort"
)

func init1()  {
	var user map[string]int = make(map[string]int)
	user["abc"] = 38
	user["user1"] = 10000
	user["user2"] = 10001

	fmt.Printf("user1: %d\n", user["user1"])
}

func init2()  {
	var m = map[string]int {
		"user01": 100,
		"user02": 1002,
	}
	fmt.Printf("user01: %d\n",m["user01"])
	fmt.Printf("user: %#v\n",m)

}

var whiteUser = map[int]bool {
	1234345: true,
	8377: true,
}

func isWhiteUser(userid int) bool {
	_, ok := whiteUser[userid]
	return ok
}

func testWhiteUser()  {
	userid := 8377
	if isWhiteUser(userid) {
		fmt.Printf("%d ,It's in white user\n", userid)
	} else {
		fmt.Println("%d ,It's not in white user\n", userid)
	}
}

func transverse()  {
	var m = map[string]int {
		"user01": 100,
		"user02": 1002,
	}

	m["user03"] = 3000

	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	
}

func testDelete() {
	var m = map[string]int {
		"user01": 100,
		"user02": 1002,
	}

	m["user03"] = 3000

	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	fmt.Println(".................")
	delete(m, "user03")

	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

}

func testMapCopy()  {
	var m = map[string]int {
		"user01": 100,
		"user02": 1002,
	}

	m["user03"] = 3000
	fmt.Printf("origin map: %#v\n", m)

	m["user02"] = 200000
	fmt.Printf("changed map: %#v\n", m)
}

func testMapSort()  {
	var m = map[string]int {
		"user01": 100,
		"user02": 1002,
	}

	m["user03"] = 3000

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("key: %s, value: %d\n", key, m[key])
	}
}

func testMapSlice()  {
	var ms []map[string]int
	ms = make([]map[string]int, 5)

	for i, v := range ms {
		ms[i] = make(map[string]int, 5)
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	ms[0]["abc"] = 100

	for k, v := range ms[0] {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}

func main() {
	//init1()
	//init2()
	//testWhiteUser()
	//transverse()
	//testDelete()
	//testMapCopy()
	//testMapSort()
	testMapSlice()
}
