package main

import (
	"fmt"
	"strings"
)

func splitString(s rune) bool {
	var m = map[string]string{
		"alpnum": "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
	}

	if strings.ContainsRune(m["alpnum"], s) {
		return false
	} else {
		return true
	}
}

func main() {
	var s = "hello world, I'm xue, how are you? hello! I'm from china. my id is 134324. stu01"

	//fmt.Println(strings.FieldsFunc(s, splitString))

	var mc = map[string]int{}
	for _, w := range strings.FieldsFunc(s, splitString) {
		if _, ok := mc[w]; ok {
			mc[w] += 1
		} else {
			mc[w] = 1
		}
	}

	//fmt.Printf("%#v\n", mc)
	for k, v := range mc {
		fmt.Printf("%v: %v\n", k, v)
	}
}
