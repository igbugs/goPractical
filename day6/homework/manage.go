package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"time"
)

var Db *sql.DB

func init() {
	dsn := "gouser:123456@tcp_chat(192.168.247.133:3306)/golang?charset=utf8mb4&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql failed, err: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("ping mysql failed, err: %v\n", err)
	}
	Db = db

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	bm := BookMgr{}
	b := Book{}
	qb := Qbook{}
	sm := StudentMgr{}
	s := Student{}

	qs := Qstu{}

	bstu := BookStu{}

	for {
		switch rt := ShowMenu(reader, 0); rt {
		case 1:
			fmt.Println("  >> 进入书籍管理系统")
			for {
				switch rt := ShowMenu(reader, 1); rt {
				case 1:
					fmt.Printf("请输入添加的书籍的Sn: ")
					input := ReadInput(reader)
					b.Sn = input
					fmt.Printf("请输入添加的书籍的Name: ")
					input = ReadInput(reader)
					b.Name = input
					fmt.Printf("请输入添加的书籍的Num: ")
					input = ReadInput(reader)
					i, _ := strconv.Atoi(input)
					b.Num = uint(i)
					fmt.Printf("请输入添加的书籍的Author: ")
					input = ReadInput(reader)
					b.Author = input
					fmt.Printf("请输入添加的书籍的PublishDate: ")
					input = ReadInput(reader)
					t, _ := time.Parse("2006-01-02 15:04:05", input)
					b.Publish = t
					bm.AddBook(b, Db)
				case 2:
					fmt.Printf("请输入删除的书籍的Sn: ")
					input := ReadInput(reader)
					b.Sn = input
					bm.DelBook(b, Db)
				case 3:
					fmt.Printf("请输入更新的书籍的Sn: ")
					input := ReadInput(reader)
					if input == "" {
						fmt.Printf("you must input the book SN.\n")
						goto FLAG
					} else {
						b.Sn = input
					}
					fmt.Printf("请输入更新的书籍的Name(or Enter): ")
					input = ReadInput(reader)
					b.Name = input
					fmt.Printf("请输入更新的书籍的Num(or Enter): ")
					input = ReadInput(reader)
					i, _ := strconv.Atoi(input)
					b.Num = uint(i)
					fmt.Printf("请输入更新的书籍的Author(or Enter): ")
					input = ReadInput(reader)
					b.Author = input
					fmt.Printf("请输入更新的书籍的PublishDate(or Enter): ")
					input = ReadInput(reader)
					t, _ := time.Parse("2006-01-02 15:04:05", input)
					b.Publish = t
					bm.UpdateBook(b, Db)
				case 4:
					fmt.Printf("请输入检索的书籍的Sn: ")
					input := ReadInput(reader)
					qb.Sn = input
					fmt.Printf("请输入检索的书籍的Name(or Enter): ")
					input = ReadInput(reader)
					qb.Name = input
					fmt.Printf("请输入检索的书籍的Num(or Enter): ")
					input = ReadInput(reader)
					qb.Num = input
					fmt.Printf("请输入检索的书籍的Author(or Enter): ")
					input = ReadInput(reader)
					qb.Author = input
					fmt.Printf("请输入检索的书籍的pulish 起始日期: ")
					input = ReadInput(reader)
					qb.BeginPubDate = input
					fmt.Printf("请输入检索的书籍的pulish 终止日期: ")
					input = ReadInput(reader)
					qb.EndPubDate = input
					r := bm.Query(qb, Db)
					fmt.Printf("查询的输出结果: \n%v\n", r)
				}

			FLAG:
				if IsExit(reader, "是否继续进行书籍管理操作") {
					continue
				} else {
					break
				}
			}
		case 2:
			fmt.Println("  >> 进入学生管理系统")
			for {
				switch rt := ShowMenu(reader, 2); rt {
				case 1:
					fmt.Printf("请输入添加的学生的CardId: ")
					input := ReadInput(reader)
					s.CardId = input
					fmt.Printf("请输入添加的学生的Name: ")
					input = ReadInput(reader)
					s.Name = input
					fmt.Printf("请输入添加的学生的Sex: ")
					input = ReadInput(reader)
					s.Sex = input
					fmt.Printf("请输入添加的学生的Grade: ")
					input = ReadInput(reader)
					s.Grade = input
					sm.AddStudent(s, Db)
				case 2:
					fmt.Printf("请输入删除的学生的CardId: ")
					input := ReadInput(reader)
					s.CardId = input
					sm.DelStudent(s, Db)
				case 3:
					fmt.Printf("请输入更新的学生的CardId: ")
					input := ReadInput(reader)
					if input == "" {
						fmt.Printf("you must input the card id.")
					} else {
						s.CardId = input
					}
					fmt.Printf("请输入更新的学生的Name(or Enter): ")
					input = ReadInput(reader)
					s.Name = input
					fmt.Printf("请输入更新的学生的Sex(or Enter): ")
					input = ReadInput(reader)
					s.Sex = input
					fmt.Printf("请输入更新的学生的Grade(or Enter): ")
					input = ReadInput(reader)
					s.Grade = input
					sm.UpdateStudent(s, Db)
				case 4:
					fmt.Printf("请输入检索的学生的CardId: ")
					input := ReadInput(reader)
					qs.CardId = input
					fmt.Printf("请输入检索的学生的Name(or Enter): ")
					input = ReadInput(reader)
					qs.Name = input
					fmt.Printf("请输入检索的学生的Sex(or Enter): ")
					input = ReadInput(reader)
					qs.Sex = input
					fmt.Printf("请输入检索的学生的Grade(or Enter): ")
					input = ReadInput(reader)
					qs.Grade = input
					r := sm.Query(qs, Db)
					fmt.Printf("查询的输出结果: \n%v\n", r)
				}

				if IsExit(reader, "是否继续进行学生管理操作") {
					continue
				} else {
					break
				}
			}
		case 3:
			fmt.Println("  >> 进入借书管理系统")
			for {
				switch rt := ShowMenu(reader, 3); rt {
				case 1:
					fmt.Printf("请输入你的学生CardId: ")
					cardid := ReadInput(reader)
					fmt.Printf("请输入你的所借书的SN: ")
					booksn := ReadInput(reader)
					bstu.BorrowBook(booksn, cardid, Db)
				case 2:
					fmt.Printf("请输入你的学生CardId: ")
					cardid := ReadInput(reader)
					fmt.Printf("请输入你的所借书的SN: ")
					booksn := ReadInput(reader)
					bstu.ReturnBook(booksn, cardid, Db)
				case 3:
					fmt.Printf("请输入你的学生CardId: ")
					cardid := ReadInput(reader)
					r := bstu.QueryBystuid(cardid, Db)
					fmt.Printf("cardid: %s 借的书籍有:\n%v\n", cardid, r)

					fmt.Printf("请输入你的所借书的SN: ")
					booksn := ReadInput(reader)
					r1 := bstu.QueryBybid(booksn, Db)
					fmt.Printf("booksn: %s 被借出给的同学有:\n%v\n", booksn, r1)
				}

				if IsExit(reader, "是否继续进行书籍管理操作") {
					continue
				} else {
					break
				}
			}
		case 4:
			fmt.Println("退出图书管理系统.")
			os.Exit(0)
		}

		if IsExit(reader, "是否进入其他管理系统或退出") {
			continue
		} else {
			break
		}
	}
}
