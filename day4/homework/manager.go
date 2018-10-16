package main

import (
	"day4/homework/util"
	"bufio"
	"os"
	"fmt"
	"day4/homework/book"
	"day4/homework/student"
	"strconv"
)

func BorrowBook(r *bufio.Reader) {
	sm := student.ReadStudentMap("C:/GoProject/Go3Project/src/day4/homework/studentMap.db")
	bm := book.ReadBookMap("C:/GoProject/Go3Project/src/day4/homework/bookMap.db")

	fmt.Printf("请输入你的学生Id: ")
	input := util.ReadInput(r)
	stuId, _ := strconv.Atoi(input)

	if _, ok := sm[stuId]; ok {
		smp := sm[stuId]
		fmt.Printf("请输入你借书的Id编号: ")
		input := util.ReadInput(r)
		bookId, _ := strconv.Atoi(input)

		if _, ok := bm[bookId]; ok {
			bmp := bm[bookId]
			if bmp.Number == 0 {
				fmt.Printf("书籍名字 %s Id %d 的书没有库存\n", bmp.Name, bmp.Id)
			} else {
				borrowbook := "(" + strconv.Itoa(bmp.Id) + ")" + bmp.Name
				borrowstudent := "(" + strconv.Itoa(smp.Id) + ")" + smp.Name
				(&smp).BorrowBook = append((&smp).BorrowBook, borrowbook)
				(&bmp).BorrowStudent = append((&bmp).BorrowStudent, borrowstudent)
				(&bmp).Number -= 1
				delete(sm, stuId)
				sm[stuId] = smp
				delete(bm, bookId)
				bm[bookId] = bmp
			}
		} else {
			fmt.Printf("你输入的书籍Id 不存在，请进行检查.\n")
		}
	} else {
		fmt.Printf("你输入的学生Id 不存在，请进行注册.\n")
	}

	util.WriteMap(sm, "C:/GoProject/Go3Project/src/day4/homework/studentMap.db")
	util.WriteMap(bm, "C:/GoProject/Go3Project/src/day4/homework/bookMap.db")
}

func ReturnBook(r *bufio.Reader) {
	sm := student.ReadStudentMap("C:/GoProject/Go3Project/src/day4/homework/studentMap.db")
	bm := book.ReadBookMap("C:/GoProject/Go3Project/src/day4/homework/bookMap.db")

	fmt.Printf("请输入你的学生Id: ")
	input := util.ReadInput(r)
	stuId, _ := strconv.Atoi(input)

	if _, ok := sm[stuId]; ok {
		smp := sm[stuId]
		fmt.Printf("你当前在借的书籍信息: \n %v\n", smp.BorrowBook)
		fmt.Printf("请输入你还书的Id编号: ")
		input := util.ReadInput(r)
		bookId, _ := strconv.Atoi(input)

		if _, ok := bm[bookId]; ok {
			bmp := bm[bookId]

			borrowbook := "(" + strconv.Itoa(bmp.Id) + ")" + bmp.Name
			borrowstudent := "(" + strconv.Itoa(smp.Id) + ")" + smp.Name

			var bs []string
			for _, s := range (&smp).BorrowBook {

				if s != borrowbook {
					bs = append(bs, s)
				}
			}
			(&smp).BorrowBook = bs

			var ss []string
			for _, s := range (&bmp).BorrowStudent {
				if s != borrowstudent {
					ss = append(ss, s)
				}
			}
			(&bmp).BorrowStudent = ss

			(&bmp).Number += 1
			delete(sm, stuId)
			sm[stuId] = smp
			delete(bm, bookId)
			bm[bookId] = bmp
		} else {
			fmt.Printf("你输入的书籍Id 不存在，请进行检查.\n")
		}
	} else {
		fmt.Printf("你输入的学生Id 不存在，请进行注册.\n")
	}

	util.WriteMap(sm, "C:/GoProject/Go3Project/src/day4/homework/studentMap.db")
	util.WriteMap(bm, "C:/GoProject/Go3Project/src/day4/homework/bookMap.db")
}

func StudentBorrowInfo(r *bufio.Reader) {
	sm := student.ReadStudentMap("C:/GoProject/Go3Project/src/day4/homework/studentMap.db")
	bm := book.ReadBookMap("C:/GoProject/Go3Project/src/day4/homework/bookMap.db")

	fmt.Printf("请输入你的学生Id: ")
	input := util.ReadInput(r)
	stuId, _ := strconv.Atoi(input)

	if _, ok := sm[stuId]; ok {
		smp := sm[stuId]
		fmt.Printf("你当前在借的书籍信息: \n %v\n", smp.BorrowBook)

		fmt.Printf("请输入你想查询和你借同一本书的Id编号: ")
		input := util.ReadInput(r)
		bookId, _ := strconv.Atoi(input)

		if _, ok := bm[bookId]; ok {
			bmp := bm[bookId]

			fmt.Printf("和你借同一本输的同学有: \n %v\n", bmp.BorrowStudent)

		} else {
			fmt.Printf("你输入的书籍Id 不存在，请进行检查.\n")
		}
	} else {
		fmt.Printf("你输入的学生Id 不存在，请进行注册.\n")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	b := book.Book{}
	s := student.Student{}

	for {
		switch rt := util.ShowMenu(reader, 0); rt {
		case 1:
			fmt.Println("  >> 进入书籍管理系统")
			for {
				switch rt := util.ShowMenu(reader, 1); rt {
				case 1:
					b.EntryBook(reader)
				case 2:
					book.DeleteBook(reader)
				case 3:
					book.ModifyBook(reader)
				case 4:
					book.QueryBook(reader)
				}

				if util.IsExit(reader, "是否继续进行书籍管理操作") {
					continue
				} else {
					break
				}
			}
		case 2:
			fmt.Println("  >> 进入学生管理系统")
			for {
				switch rt := util.ShowMenu(reader, 2); rt {
				case 1:
					s.EntryStudent(reader)
				case 2:
					student.DeleteStudent(reader)
				case 3:
					student.ModifyStudent(reader)
				case 4:
					student.QueryStudent(reader)
				}

				if util.IsExit(reader, "是否继续进行学生管理操作") {
					continue
				} else {
					break
				}
			}
		case 3:
			fmt.Println("  >> 进入借书管理系统")
			for {
				switch rt := util.ShowMenu(reader, 3); rt {
				case 1:
					BorrowBook(reader)
				case 2:
					ReturnBook(reader)
				case 3:
					StudentBorrowInfo(reader)
				}

				if util.IsExit(reader, "是否继续进行书籍管理操作") {
					continue
				} else {
					break
				}
			}
		case 4:
			fmt.Println("退出图书管理系统.")
			os.Exit(0)
		}

		if util.IsExit(reader, "是否进入其他管理系统或退出") {
			continue
		} else {
			break
		}
	}

}
