package main

import (
	"fmt"
)

func charCount(s string) map[string]int {
	var chars = []rune(s)
	m := make(map[string]int)

	for _, c := range chars {
		if _, ok := m[string(c)]; ok {
			m[string(c)] += 1
		} else {
			m[string(c)] = 1
		}
	}
	return m
}

func main() {
	var s = "我在学习GO, nizaiganma 234o 我的god, ! 2%^&$%^"
	for k, v := range charCount(s) {
		fmt.Printf("m[%s] = %d\n", k, v)
	}
}
