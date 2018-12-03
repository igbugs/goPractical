package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	numStr     = "0123456789"
	charStr    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	specialStr = "~!@#$%^&*()_+{}[]|:;<>,?/"
)

func getArgs() {
	flag.IntVar(&length, "l", 16, "-l 密码的长度(length)")
	flag.StringVar(&charset, "t", "mix",
		`-t 指定密码生成的字符集, 
num: 密码只有数字,
char: 密码只有字母,
mix: 密码包含字母与数字,
advance: 密码含有数字、字目和特殊字符`)
	flag.Parse()
}

func generatePasswd() string {
	var passwd = make([]byte, length, length)
	var selectCharSet string

	switch charset {
	case "num":
		selectCharSet = numStr
	case "char":
		selectCharSet = charStr
	case "mix":
		selectCharSet = numStr + charStr
	case "advance":
		selectCharSet = numStr + charStr + specialStr
	default:
		selectCharSet = numStr + charStr
	}

	fmt.Printf("select char set: %s\n", selectCharSet)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		passwd[i] = selectCharSet[rand.Intn(len(selectCharSet))]
	}

	return string(passwd)
}

func main() {
	getArgs()
	fmt.Printf("length: %d, charset: %s\n", length, charset)

	passwd := generatePasswd()
	fmt.Println(passwd)
}
