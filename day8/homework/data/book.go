package data

import (
	"database/sql"
	"day8/homework/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func InsertBook(book *models.Book) (err error) {
	if book == nil {
		err = fmt.Errorf("invalid book parameter")
		return
	}

	sqlstr := "select book_id from book where book_id=?"
	var bookId string
	err = Db.Get(&bookId, sqlstr, book.BookId)
	if err == sql.ErrNoRows {
		// 插入操作
		sqlstr = `insert into book (
					  book_id, author, book_name, publish_time, stock_num, create_time
				  ) values (?, ?, ?, ?, ?, ?)`
		_, err = Db.Exec(sqlstr, book.BookId, book.Author, book.BookName, book.PublishTime, book.StockNum, time.Now())
		if err != nil {
			return
		}
		return
	}

	if err != nil {
		return
	}

	err = fmt.Errorf("book_id:%s is already exists", bookId)
	return
}

func UpdateBook(book *models.Book) (err error) {
	if book == nil {
		err = fmt.Errorf("invalid book parameter")
		return
	}

	// 更新操作
	sqlstr := `update book set 
					author = ?, book_name = ?, publish_time = ?, stock_num = stock_num + ?, update_time = ?
				where 
				    book_id = ?`
	result, err := Db.Exec(sqlstr, book.Author, book.BookName, book.PublishTime, book.StockNum, time.Now(), book.BookId)
	if err != nil {
		return
	}

	affects, err := result.RowsAffected()
	if err != nil {
		return
	}

	if affects == 0 {
		err = fmt.Errorf("update book failed, book_id:%s, not found", book.BookId)
		return
	}
	return
}

func QueryBook(bookId string) (book *models.Book, err error) {
	// 查询操作
	sqlstr := `select
					author, book_name, publish_time, stock_num, book_id, id
				from
					book
				where
					book_id = ?`
	book = &models.Book{}
	err = Db.Get(book, sqlstr, bookId)
	if err != nil {
		return
	}
	return
}

func DeleteBook(bookid string) (err error) {
	sqlstr := "select book_id from book where book_id=?"
	var bookId string
	err = Db.Get(&bookId, sqlstr, bookid)
	if err == sql.ErrNoRows {
		fmt.Printf("bookid: %s not found", bookid)
		return err
	}

	if err != nil {
		return err
	}

	// 删除操作
	sqlstr = `delete from
					book
				where 
					book_id = ?`
	_, err = Db.Exec(sqlstr, bookid)
	if err != nil {
		return nil
	}
	return
}
