package main

import (
	"fmt"
	"database/sql"
	"time"
	"strconv"
	"strings"
)

type Book struct {
	ID int				`db:"id"`
	Sn string			`db:"sn"`
	Name string			`db:"name"`
	Num uint			`db:"number"`
	Author string		`db:"author"`
	Publish time.Time	`db:"publish_date"`
}

type BookMgr struct {}

type Qbook struct {
	ID string
	Sn string
	Name string
	Num string
	Author string
	BeginPubDate string
	EndPubDate string
}

func (bm BookMgr) Query(qb Qbook, db *sql.DB) []Book {
	sqlStr := "SELECT * FROM book WHERE 1=1"

	if qb.ID != "" {
		sqlStr += " AND id=" +  qb.ID
	}
	if qb.Sn != "" {
		sqlStr += " AND sn='" + qb.Sn + "'"
	}
	if qb.Name != "" {
		sqlStr += " AND name='" + qb.Name + "'"
	}
	if qb.Num != "" {
		sqlStr += " AND number='" + qb.Num + "'"
	}
	if qb.Author != "" {
		sqlStr += " AND author='" + qb.Author + "'"
	}
	if qb.BeginPubDate != "" && qb.EndPubDate != "" {
		sqlStr += " AND publish_date BETWEEN '" + qb.BeginPubDate + "'" + " AND '" + qb.EndPubDate + "'"
	}

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
	}

	rows, err := stmt.Query()
	defer func() {
		if rows != nil {
			rows.Close()
		}
		if stmt != nil {
			stmt.Close()
		}
	}()

	if err == sql.ErrNoRows {
		fmt.Printf("not found data of id %d\n", 1)
	}

	if err != nil {
		fmt.Printf("scan row failed! err: %v\n", err)
	}

	var bs []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Sn, &book.Name, &book.Num, &book.Author, &book.Publish)
		if err != nil {
			fmt.Printf("scan row failed! err: %v\n", err)
		}
		bs = append(bs, book)
	}
	return bs
}

func (bm BookMgr) AddBook(b Book, db *sql.DB) {
	var qb Qbook
	if b.Sn != "" {
		qb.Sn = b.Sn
	}

	r := bm.Query(qb, db)
	if len(r) == 0 {
		sqlStr := "INSERT INTO book (sn, name, number, author, publish_date) VALUE (?, ?, ?, ?, ?)"
		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			fmt.Printf("prepare failed, err:%v\n", err)
			return
		}
		defer func() {
			if stmt != nil {
				stmt.Close()
			}
		}()

		result, err := stmt.Exec(b.Sn, b.Name, b.Num, b.Author, b.Publish)
		if err != nil {
			fmt.Printf("exec failed! err: %v\n", err)
			return
		}

		id, _ := result.LastInsertId()
		affectRows, _ := result.RowsAffected()
		fmt.Printf("last insert id:%d affect rows:%d\n", id, affectRows)
	} else {
		fmt.Printf("book Sn: %s is exist!", b.Sn)
		return
	}
}

func (bm BookMgr) DelBook(b Book, db *sql.DB) {
	var qb Qbook
	if b.Sn != "" {
		qb.Sn = b.Sn
	}

	r := bm.Query(qb, db)
	if len(r) != 0 {
		sqlStr := "DELETE FROM book WHERE sn=?"
		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			fmt.Printf("prepare failed, err:%v\n", err)
			return
		}
		defer func() {
			if stmt != nil {
				stmt.Close()
			}
		}()

		result, err := stmt.Exec(b.Sn)
		if err != nil {
			fmt.Printf("exec failed! err: %v\n", err)
			return
		}

		affectRows, _ := result.RowsAffected()
		fmt.Printf("affect rows:%d\n", affectRows)
	} else {
		fmt.Printf("book Sn: %s is exist!", b.Sn)
		return
	}
}

func (bm BookMgr) UpdateBook(b Book, db *sql.DB) {
	var qb Qbook
	if b.Sn != "" {
		qb.Sn = b.Sn
	}

	r := bm.Query(qb, db)
	if len(r) != 0 {
		sqlStr := "UPDATE book SET"
		rb := r[0]
		if rb.Name != b.Name && b.Name != "" {
			sqlStr += " name='" + b.Name + "',"
		}
		if rb.Num != b.Num && b.Num != 0 {
			sqlStr += " number='" + strconv.Itoa(int(b.Num)) + "',"
		}
		if rb.Author != b.Author && b.Author != "" {
			sqlStr += " author='" + b.Author + "',"
		}
		if !b.Publish.IsZero() {
			sqlStr += " publish_date='" + b.Publish.Format("2006-01-02 15:04:05") + "',"
		}

		sqlStr = strings.Trim(sqlStr, ",") + " WHERE sn='" + b.Sn + "'"
		fmt.Println(sqlStr)
		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			fmt.Printf("prepare failed, err:%v\n", err)
			return
		}
		defer func() {
			if stmt != nil {
				stmt.Close()
			}
		}()

		result, err := stmt.Exec()
		if err != nil {
			fmt.Printf("exec failed! err: %v\n", err)
			return
		}

		affectRows, _ := result.RowsAffected()
		fmt.Printf("affect rows:%d\n", affectRows)
	} else {
		fmt.Printf("book Sn: %s is exist!", b.Sn)
		return
	}
}
