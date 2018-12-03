package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInput(r *bufio.Reader) string {
	input, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	spaceDel := strings.Replace(input, " ", "", -1)
	return strings.Replace(spaceDel, "\n", "", -1)
}

func WriteMap(m interface{}, dbname string) {
	if m1, err := json.Marshal(m); err == nil {
		ioutil.WriteFile(dbname, m1, 0644)
	}
}

func ShowMenu(r *bufio.Reader, sub int) int {
	switch sub {
	case 0:
		fmt.Printf("欢迎进入图书管理系统，请进行功能选择: \n")
		fmt.Printf("\t1. 书籍管理\n")
		fmt.Printf("\t\t进行书籍的信息录入、删除、更改、检索的操作\n")
		fmt.Printf("\t2. 学生管理\n")
		fmt.Printf("\t\t进行学生信息的录入、删除、更改、查询的操作\n")
		fmt.Printf("\t3. 借书管理\n")
		fmt.Printf("\t\t进行书籍的借出、学生的还书、学生的借书信息查询\n")
		fmt.Printf("\t4. 退出系统\n")

		fmt.Printf("请输入选择: ")
		input := ReadInput(r)
		i, _ := strconv.Atoi(input)
		return i
	case 1:
		fmt.Println("  >> 进入书籍管理子菜单")
		fmt.Printf("\t\t1. 书籍录入\n")
		fmt.Printf("\t\t2. 书籍删除\n")
		fmt.Printf("\t\t3. 书籍更改\n")
		fmt.Printf("\t\t4. 书籍检索\n")

		fmt.Printf("  >> 请输入选择: ")
		input := ReadInput(r)
		i, _ := strconv.Atoi(input)
		return i
	case 2:
		fmt.Println("  >> 进入学生管理子菜单")
		fmt.Printf("\t\t1. 学生录入\n")
		fmt.Printf("\t\t2. 学生删除\n")
		fmt.Printf("\t\t3. 学生更改\n")
		fmt.Printf("\t\t4. 学生信息查询\n")

		fmt.Printf("  >> 请输入选择: ")
		input := ReadInput(r)
		i, _ := strconv.Atoi(input)
		return i
	case 3:
		fmt.Println("  >> 进入借书管理子菜单")
		fmt.Printf("\t\t1. 借书\n")
		fmt.Printf("\t\t2. 还书\n")
		fmt.Printf("\t\t3. 学生的借书信息查询\n")

		fmt.Printf("  >> 请输入选择: ")
		input := ReadInput(r)
		i, _ := strconv.Atoi(input)
		return i
	default:
		return 0
	}
}

func IsExit(r *bufio.Reader, s string) bool {
	var f bool
	fmt.Printf("%s (yes|no)? ", s)
	input := ReadInput(r)
	exit := input

	switch exit {
	case "yes", "y", "YES", "Y":
		f = true
	case "no", "n", "NO", "N":
		f = false
	default:
		f = false
	}
	return f
}
