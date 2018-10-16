package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"strings"
	"errors"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

type Student struct {
	Username string
	Score    float32
	Grade    string
	Sex      string
}

type studentMap map[string]Student

var (
	//stuMap = make(map[string]Student)
	stuMap = studentMap{}
	opType string
)

func readStudentMap() {
	if file, err := os.Open("./studentmap.db"); err == nil {
		defer file.Close()
		if contents, err := ioutil.ReadAll(file); err == nil {
			json.Unmarshal(contents, &stuMap)
		}
	}
}

func writeStudentMap() {
	if m, err := json.Marshal(stuMap); err == nil {
		ioutil.WriteFile("./studentmap.db", m, 0644)
	}
}

func addStudent(username string, score float32, grade, sex string) {
	if _, ok := stuMap[username]; !ok {
		stuMap[username] = Student{
			Username: username,
			Score:    score,
			Grade:    grade,
			Sex:      sex,
		}
		writeStudentMap()
		fmt.Printf("%v\n", stuMap[username])
	} else {
		fmt.Printf("学生 %s 已经存在。", username)
	}
}

func modifyStudent(username string, field string, stu *Student, reader *bufio.Reader) {
	if _, ok := stuMap[username]; ok {
		var user = *stu
		switch field {
		case "Username":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input := readInput(reader)
			user.Username = input
			deleteStudent(username)
			stuMap[input] = user
		case "Score":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input := readInput(reader)
			s, _ := strconv.ParseFloat(input, 32)
			user.Score = float32(s)
			deleteStudent(username)
			stuMap[username] = user
		case "Grade":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input := readInput(reader)
			user.Grade = input
			deleteStudent(username)
			stuMap[username] = user
		case "Sex":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input := readInput(reader)
			user.Sex = input
			deleteStudent(username)
			stuMap[username] = user
		}
		writeStudentMap()
	} else {
		fmt.Printf("学生 %s 不存在。", username)
	}
}

func deleteStudent(username string) {
	if _, ok := stuMap[username]; ok {
		delete(stuMap, username)
		writeStudentMap()
	} else {
		fmt.Printf("学生 %s 不存在。", username)
	}
}

func showStudent(username string) error {
	if username == "all" {
		if len(stuMap) == 0 {
			fmt.Println("还没有录入学生的信息！")
			return errors.New("empty dict in stuMap")
		} else {
			fmt.Println("所有的学生的信息:")
			for k, v := range stuMap {
				fmt.Printf("  %s:  %v\n", k, v)
			}
		}
	} else {
		if _, ok := stuMap[username]; ok {
			fmt.Printf("%s 学生的信息为: \n%v\n", username, stuMap[username])
		} else {
			fmt.Printf("学生 %s 不存在. ", username)
			return errors.New("student isn't exist")
		}
	}
	return nil
}

func getArgs() {
	flag.StringVar(&opType, "m", "",
		`--manager 指定操作的类型,
	  add 添加学生信息,
	  modify 更改学生的信息,
	  show 打印所有学生的信息`)
	flag.Parse()
}

func isExit(exit string) bool {
	var f bool
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

func readInput(r *bufio.Reader) string {
	input, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		//return err
	}
	return strings.Replace(input, "\n", "", -1)
}

func main() {
	getArgs()
	var (
		username string
		score    float32
		grade    string
		sex      string
		exit     string
		stuField string
	)

	reader := bufio.NewReader(os.Stdin)
	switch opType {
	case "add":
		for {
			readStudentMap()
			fmt.Printf("请输入添加的学生的名字: ")
			input := readInput(reader)
			username = input

			if _, ok := stuMap[username]; ok {
				fmt.Printf("%s 学生已经存在, 只能进行更改. ", username)
				break
			}

			fmt.Printf("请输入添加的学生的分数: ")
			input = readInput(reader)
			s, _ := strconv.ParseFloat(input, 32)
			score = float32(s)

			fmt.Printf("请输入添加的学生的年级: ")
			input = readInput(reader)
			grade = input

			fmt.Printf("请输入添加的学生的性别: ")
			input = readInput(reader)
			sex = input

			addStudent(username, score, grade, sex)

			fmt.Printf("是否继续添加学生信息(yes|no)? ")
			input = readInput(reader)
			exit = input

			if isExit(exit) {
				continue
			} else {
				break
			}
		}
	case "modify":
		for {
			readStudentMap()

			fmt.Printf("请输入更改的学生的名字: ")
			input := readInput(reader)
			username = input

			if _, ok := stuMap[username]; !ok {
				fmt.Printf("%s 学生不存在！ ", username)
				break
			}

			fmt.Printf("请输入更改的学生的信息(Username|Score|Grade|Sex): ")
			input = readInput(reader)
			stuField = input
			var stu = stuMap[username]
			var user *Student = &stu
			modifyStudent(username, stuField, user, reader)

			fmt.Printf("是否继续更改 %s 的信息(yes|no)? ", username)
			input = readInput(reader)
			exit = input

			if isExit(exit) {
				continue
			} else {
				break
			}
		}
	case "delete":
		for {
			readStudentMap()

			fmt.Printf("请输入删除的学生的名字: ")
			input := readInput(reader)
			username = input

			if _, ok := stuMap[username]; !ok {
				fmt.Printf("%s 学生不存在！ ", username)
				break
			}

			deleteStudent(username)

			fmt.Printf("是否继续删除操作 (yes|no)? ")
			input = readInput(reader)
			exit = input

			if isExit(exit) {
				continue
			} else {
				break
			}
		}
	case "show":
		for {
			readStudentMap()

			fmt.Printf("请输入需要显示的学生名字(username|all): ")
			input := readInput(reader)
			username = input

			if err := showStudent(username); err != nil {
				fmt.Println(err)
				break
			}

			fmt.Printf("是否继续显示学生信息 (yes|no)? ")
			input = readInput(reader)
			exit = input

			if isExit(exit) {
				continue
			} else {
				break
			}
		}
	default:
		fmt.Println(`Usage: ./4_student_manager -m, --manager [add|modify|delete|show]
			   -h, --help  (help info)`)
	}
}
