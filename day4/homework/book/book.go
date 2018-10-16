package book

import (
	"bufio"
	"fmt"
	"strconv"
	"day4/homework/util"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Book struct {
	Id            int
	Name          string
	Number        int
	Author        string
	PubDate       string
	BorrowStudent []string
}

type BooksMap map[int]Book

func ReadBookMap(dbname string) BooksMap {
	var bm = BooksMap{}
	if file, err := os.Open(dbname); err == nil {
		defer file.Close()
		if contents, err := ioutil.ReadAll(file); err == nil {
			json.Unmarshal(contents, &bm)
		}
	}
	return bm
}

func (b *Book) EntryBook(r *bufio.Reader) {
	bm := ReadBookMap("C:/GoProject/Go3Project/src/day4/homework/bookMap.db")

	fmt.Printf("请输入添加的书籍的Id: ")
	input := util.ReadInput(r)
	for _, s := range input {
		if s >= '0' && s <= '9' {
			b.Id, _ = strconv.Atoi(input)

			if _, ok := bm[b.Id]; !ok {
				fmt.Printf("请输入添加的书籍的名字: ")
				input = util.ReadInput(r)
				b.Name = input

				fmt.Printf("请输入添加的书籍的数量: ")
				input = util.ReadInput(r)
				b.Number, _ = strconv.Atoi(input)

				fmt.Printf("请输入添加的书籍的作者: ")
				input = util.ReadInput(r)
				b.Author = input

				fmt.Printf("请输入添加的书籍的出版日期: ")
				input = util.ReadInput(r)
				b.PubDate = input

				fmt.Printf("新添加的书籍%s信息为: %v\n", b.Name, *b)
			} else {
				fmt.Printf("book id %d exist, please check it!\n", b.Id)
			}
		} else {
			fmt.Println("ERROR: 输入的不为纯数字.")
			break
		}
	}

	bm[b.Id] = *b
	util.WriteMap(bm, "C:/GoProject/Go3Project/src/day4/homework/bookMap.db")
	//return bm
}

func DeleteBook(r *bufio.Reader) {
	bm := ReadBookMap("C:/GoProject/Go3Project/src/day4/homework/bookMap.db")

	fmt.Printf("请输入删除书籍的Id编号: ")
	input := util.ReadInput(r)
	id, _ := strconv.Atoi(input)
	delete(bm, id)
	util.WriteMap(bm, "C:/GoProject/Go3Project/src/day4/homework/bookMap.db")
}

func ModifyBook(r *bufio.Reader) {
	bm := ReadBookMap("C:/GoProject/Go3Project/src/day4/homework/bookMap.db")

	fmt.Printf("请输入需要更改书籍的Id编号: ")
	input := util.ReadInput(r)
	id, _ := strconv.Atoi(input)

	if _, ok := bm[id]; ok {
		fmt.Printf("请输入修改书籍的信息(Id|Name|Number|Author|PubDate): ")
		field := util.ReadInput(r)

		bmp := bm[id]
		switch field {
		case "Id":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input := util.ReadInput(r)
			(&bmp).Id, _ = strconv.Atoi(input)
		case "Name":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input := util.ReadInput(r)
			(&bmp).Name = input
		case "Number":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input := util.ReadInput(r)
			(&bmp).Number, _ = strconv.Atoi(input)
		case "Author":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input := util.ReadInput(r)
			(&bmp).Author = input
		case "PubDate":
			fmt.Printf("请输入 %s 的更改值: ", field)
			input := util.ReadInput(r)
			(&bmp).PubDate = input
		}
		delete(bm, id)
		bm[id] = bmp
		util.WriteMap(bm, "C:/GoProject/Go3Project/src/day4/homework/bookMap.db")
	} else {
		fmt.Printf("输入的书籍Id %d 不存在。", id)
	}
}

func QueryBook(r *bufio.Reader) {
	bm := ReadBookMap("C:/GoProject/Go3Project/src/day4/homework/bookMap.db")

	var condition []map[string]string
	for {
		fmt.Printf("请输入检索书籍的条件(Id|Name|Number|Author|PubDate) (default: all book show): ")
		inputK := util.ReadInput(r)
		if inputK != "" {
			fmt.Printf("请输入检索书籍的条件 %s 的值: ", inputK)
			inputV := util.ReadInput(r)
			m := map[string]string{inputK: inputV}
			condition = append(condition, m)
		} else {
			for _, bmv := range bm {
				fmt.Println(bmv)
			}
		}

		if util.IsExit(r, "是否选择其他的检索条件") {
			continue
		} else {
			break
		}
	}

	//f := func(k, v string, bm BooksMap) {
	//	for bmk, bmv := range bm {
	//		//rt := reflect.TypeOf(bmv)
	//		rv := reflect.ValueOf(bmv)
	//		switch k {
	//		case "Id":
	//			if val, _ := strconv.Atoi(v); rv.FieldByName("Id").Interface() == val {
	//				tmpMap[bmk] = bmv
	//			} else {
	//				delete(tmpMap, bmk)
	//			}
	//		case "Number":
	//			if val, _ := strconv.Atoi(v); rv.FieldByName("Number").Interface() == val {
	//				tmpMap[bmk] = bmv
	//			} else {
	//				delete(tmpMap, bmk)
	//			}
	//		case "Name":
	//			if rv.FieldByName("Name").Interface() == v {
	//				tmpMap[bmk] = bmv
	//			} else {
	//				delete(tmpMap, bmk)
	//			}
	//		case "Author":
	//			if rv.FieldByName("Author").Interface() == v {
	//				tmpMap[bmk] = bmv
	//			} else {
	//				delete(tmpMap, bmk)
	//			}
	//		case "PubDate":
	//			if rv.FieldByName("PubDate").Interface() == v {
	//				tmpMap[bmk] = bmv
	//			} else {
	//				delete(tmpMap, bmk)
	//			}
	//		}

	//for i := 0; i < rt.NumField(); i++ {
	//	if rt.Field(i).Name == "Id" || rt.Field(i).Name == "Number" {
	//		if rv.Field(i).Interface() == strconv.Atoi(v) {
	//			tmpMap[bmk] = bmv
	//		}
	//	}
	//}

	var tmpMap = BooksMap{}
	f := func(k, v string, bm BooksMap) {
		for bmk, bmv := range bm {
			switch k {
			case "Id":
				if val, _ := strconv.Atoi(v); bmv.Id == val {
					tmpMap[bmk] = bmv
				} else {
					delete(tmpMap, bmk)
				}
			case "Number":
				if val, _ := strconv.Atoi(v); bmv.Number == val {
					tmpMap[bmk] = bmv
				} else {
					delete(tmpMap, bmk)
				}
			case "Name":
				if bmv.Name == v {
					tmpMap[bmk] = bmv
				} else {
					delete(tmpMap, bmk)
				}
			case "Author":
				if bmv.Author == v {
					tmpMap[bmk] = bmv
				} else {
					delete(tmpMap, bmk)
				}
			case "PubDate":
				if bmv.PubDate == v {
					tmpMap[bmk] = bmv
				} else {
					delete(tmpMap, bmk)
				}
			}
		}
	}

	for _, sv := range condition { // condition 为map 的切片
		first := true
		for k, v := range sv { // 切片的检索条件 key 与 value 的输入值
			if len(tmpMap) == 0 && first {
				f(k, v, bm)
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
