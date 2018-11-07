package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "xyb",
		"course":  "goland",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)  // m2 == empty map
	var m3 map[string]int       // m3 == nil

	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	//for k := range m {
	//	fmt.Println(k)
	//}
	//
	//for _, v := range m {
	//	fmt.Println(v)
	//}

	fmt.Println("Getting map values")

	fmt.Println("Wrong key name")
	causeName, ok := m["cause"]     // map 中不存在的值，也可以取到
	fmt.Println(causeName)

	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if causeName, ok := m["cause"]; ok {    // 使用 ok 变量接收，key值存在与否
		fmt.Println(causeName, ok)
	} else {
		fmt.Println("key does not values")
	}

	fmt.Println("Deleting values")

	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name, ok = m["name"]   // 此处不能再使用 := 因为之前已经定义过了
	fmt.Println(name, ok)
}
