package data

import (
	"day8/homework/models"
	"testing"
	"time"
)

func init() {
	dns := "root:123456@tcp(192.168.247.133:3306)/library_mgr?parseTime=True"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertBook(t *testing.T) {
	var tests = []*models.Book{
		&models.Book{
			Author:      "jim",
			BookId:      "111111111110",
			BookName:    "C语言从入门到放弃",
			PublishTime: time.Now(),
			StockNum:    10,
		},
		&models.Book{
			Author:      "jim",
			BookId:      "111111111111",
			BookName:    "D语言从入门到放弃",
			PublishTime: time.Now(),
			StockNum:    10,
		},
		&models.Book{
			Author:      "jim",
			BookId:      "111111111112",
			BookName:    "E语言从入门到放弃",
			PublishTime: time.Now(),
			StockNum:    10,
		},
		&models.Book{
			Author:      "jim",
			BookId:      "111111111113",
			BookName:    "F语言从入门到放弃",
			PublishTime: time.Now(),
			StockNum:    10,
		},
	}

	for _, tt := range tests {
		err := InsertBook(tt)
		if err != nil {
			t.Errorf("insert book failed, err:%v", err)
			return
		}
		t.Logf("insert book succ")
	}
}

func TestUpdateBook(t *testing.T) {
	var book = models.Book{
		Author:      "jim",
		BookId:      "83883488344",
		BookName:    "C语言从入门到放弃",
		PublishTime: time.Now(),
		StockNum:    10,
	}

	err := UpdateBook(&book)
	if err != nil {
		t.Errorf("update book failed, err:%v", err)
		return
	}

	t.Logf("update book succ")
}

func TestQueryBook(t *testing.T) {
	var book *models.Book
	bookId := "83883488344"
	book, err := QueryBook(bookId)
	if err != nil {
		t.Errorf("query book failed, err:%v", err)
		return
	}

	t.Logf("query book succ, book:%#v", book)
}

func TestDeleteBook(t *testing.T) {
	bookId := "83883488344"
	err := DeleteBook(bookId)
	if err != nil {
		t.Errorf("delete book failed, err:%v", err)
		return
	}

	t.Logf("delete book succ")
}
