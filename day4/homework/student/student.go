package student

import (
	"bufio"
	"day4/homework/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Student struct {
	Id         int
	Name       string
	Grade      string
	Sex        string
	CardId     string
	BorrowBook []string
}

type StudentsMap map[int]Student

func ReadStudentMap(dbname string) StudentsMap {
	var sm = StudentsMap{}
	if file, err := os.Open(dbname); err == nil {
		defer file.Close()
		if contents, err := ioutil.ReadAll(file); err == nil {
			json.Unmarshal(contents, &sm)
		}
	}
	return sm
}

func (s *Student) EntryStudent(r *bufio.Reader) {
	sm := ReadStudentMap("C:/GoProject/Go3Project/src/day4/homework/studentMap.db")

	fmt.Printf("请输入添加的学生的id: ")
	input := util.ReadInput(r)
	for _, char := range input {
		if char >= '0' && char <= '9' {
			s.Id, _ = strconv.Atoi(input)

			if _, ok := sm[s.Id]; !ok {
				fmt.Printf("请输入添加的学生的名字: ")
				input = util.ReadInput(r)
				s.Name = input

				fmt.Printf("请输入添加的学生的班级: ")
				input = util.ReadInput(r)
				s.Grade = input

				fmt.Printf("请输入添加的学生的性别: ")
				input = util.ReadInput(r)
				s.Sex = input

				fmt.Printf("请输入添加的学生的身份证: ")
				input = util.ReadInput(r)
				s.CardId = input

				fmt.Printf("新添加的学生%s信息为: %v\n", s.Name, *s)
			} else {
				fmt.Printf("student id %d exist, please check it!\n", s.Id)
			}
		} else {
			fmt.Println("ERROR: 输入的不为纯数字.")
			break
		}
	}

	sm[s.Id] = *s
	util.WriteMap(sm, "C:/GoProject/Go3Project/src/day4/homework/studentMap.db")
	//return sm
}

func DeleteStudent(r *bufio.Reader) {
	sm := ReadStudentMap("C:/GoProject/Go3Project/src/day4/homework/studentMap.db")

	fmt.Printf("请输入删除学生的Id编号: ")
	input := util.ReadInput(r)
	id, _ := strconv.Atoi(input)
	delete(sm, id)
	util.WriteMap(sm, "C:/GoProject/Go3Project/src/day4/homework/studentMap.db")
}

func ModifyStudent(r *bufio.Reader) {
	sm := ReadStudentMap("C:/GoProject/Go3Project/src/day4/homework/studentMap.db")

	fmt.Printf("请输入需要更改学生的Id编号: ")
	input := util.ReadInput(r)
	id, _ := strconv.Atoi(input)

	if _, ok := sm[id]; ok {
		fmt.Printf("请输入修改学生的信息(Id|Name|Grade|Sex|CardId): ")
		field := util.ReadInput(r)

		smp := sm[id]
		switch field {
		case "Id":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input = util.ReadInput(r)
			(&smp).Id, _ = strconv.Atoi(input)
		case "Name":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input = util.ReadInput(r)
			(&smp).Name = input
		case "Grade":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input = util.ReadInput(r)
			(&smp).Grade = input
		case "Sex":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input = util.ReadInput(r)
			(&smp).Sex = input
		case "CardId":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input = util.ReadInput(r)
			(&smp).CardId = input
		}
		delete(sm, id)
		sm[id] = smp
		util.WriteMap(sm, "C:/GoProject/Go3Project/src/day4/homework/studentMap.db")
	} else {
		fmt.Printf("输入的书籍Id %d 不存在。", id)
	}
}

func QueryStudent(r *bufio.Reader) {
	sm := ReadStudentMap("C:/GoProject/Go3Project/src/day4/homework/studentMap.db")

	var condition []map[string]string
	for {
		fmt.Printf("请输入检索学生的条件(Id|Name|Grade|Sex|CardId) (default: all book show): ")
		inputK := util.ReadInput(r)

		if inputK != "" {
			fmt.Printf("请输入检索学生的条件 %s 的值: ", inputK)
			inputV := util.ReadInput(r)
			m := map[string]string{inputK: inputV}
			condition = append(condition, m)
		} else {
			for _, smv := range sm {
				fmt.Println(smv)
			}
		}

		if util.IsExit(r, "是否选择其他的检索条件") {
			continue
		} else {
			break
		}
	}

	var tmpMap = StudentsMap{}
	f := func(k, v string, sm StudentsMap) {
		for smk, smv := range sm {
			switch k {
			case "Id":
				if val, _ := strconv.Atoi(v); smv.Id == val {
					tmpMap[smk] = smv
				} else {
					delete(tmpMap, smk)
				}
			case "Name":
				if smv.Name == v {
					tmpMap[smk] = smv
				} else {
					delete(tmpMap, smk)
				}
			case "Grade":
				if smv.Grade == v {
					tmpMap[smk] = smv
				} else {
					delete(tmpMap, smk)
				}
			case "Sex":
				if smv.Sex == v {
					tmpMap[smk] = smv
				} else {
					delete(tmpMap, smk)
				}
			case "CardId":
				if smv.CardId == v {
					tmpMap[smk] = smv
				} else {
					delete(tmpMap, smk)
				}
			}
		}
	}

	for _, sv := range condition { // condition 为map 的切片
		first := true
		for k, v := range sv { // 切片的检索条件 key 与 value 的输入值
			if len(tmpMap) == 0 && first {
				f(k, v, sm)
				first = false
			} else {
				f(k, v, tmpMap)
			}
		}
	}

	fmt.Printf("检索条件 %v 的输出为: \n", condition)
	for _, v := range tmpMap {
		fmt.Printf("%#v\n", v)
	}
}
