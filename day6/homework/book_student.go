package main

import (
	"database/sql"
	"fmt"
)

type BookStu struct {
	Id        string `db:"id"`
	BookId    string `db:"book_id"`
	StudentId string `db:"student_id"`
}

func (bs BookStu) QueryBybid(id string, db *sql.DB) []Student {
	bm := BookMgr{}
	qb := Qbook{Sn: id}
	r := bm.Query(qb, db)

	sqlStr := "SELECT * FROM book_student WHERE book_id=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
	}

	rows, err := stmt.Query(r[0].ID)
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

	var bss []Student
	for rows.Next() {
		var b BookStu
		err := rows.Scan(&b.Id, &b.BookId, &b.StudentId)
		if err != nil {
			fmt.Printf("scan row failed! err: %v\n", err)
		}

		sm := StudentMgr{}
		qs := Qstu{ID: b.StudentId}
		r := sm.Query(qs, db)
		bss = append(bss, r[0])
	}
	return bss
}

func (bs BookStu) QueryBystuid(id string, db *sql.DB) []Book {
	sm := StudentMgr{}
	qs := Qstu{CardId: id}
	r := sm.Query(qs, db)

	sqlStr := "SELECT * FROM book_student WHERE student_id=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
	}

	rows, err := stmt.Query(r[0].ID)
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

	var bss []Book
	for rows.Next() {
		var b BookStu
		err := rows.Scan(&b.Id, &b.BookId, &b.StudentId)
		if err != nil {
			fmt.Printf("scan row failed! err: %v\n", err)
		}

		bm := BookMgr{}
		qb := Qbook{ID: b.BookId}
		r := bm.Query(qb, db)
		bss = append(bss, r[0])
	}
	return bss
}

func (bs BookStu) BorrowBook(sn string, card_id string, db *sql.DB) {
	bm := BookMgr{}
	qb := Qbook{Sn: sn}
	r := bm.Query(qb, db)
	if len(r) == 0 {
		fmt.Printf("book sn: %s isn't exists.", sn)
	} else {
		r[0].Num -= 1
		bm.UpdateBook(r[0], Db)

		sqlStr := "INSERT INTO book_student (book_id, student_id) VALUE (?, ?)"
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

		sm := StudentMgr{}
		qs := Qstu{CardId: card_id}
		r1 := sm.Query(qs, db)
		result, err := stmt.Exec(r[0].ID, r1[0].ID)
		if err != nil {
			fmt.Printf("exec failed! err: %v\n", err)
			return
		}

		id, _ := result.LastInsertId()
		affectRows, _ := result.RowsAffected()
		fmt.Printf("last insert id:%d affect rows:%d\n", id, affectRows)
	}
}

func (bs BookStu) ReturnBook(sn string, card_id string, db *sql.DB) {
	bm := BookMgr{}
	qb := Qbook{Sn: sn}
	r := bm.Query(qb, db)
	if len(r) == 0 {
		fmt.Printf("book sn: %s isn't exists.", sn)
	} else {
		r[0].Num += 1
		bm.UpdateBook(r[0], Db)

		sqlStr := "DELETE FROM book_student WHERE book_id = ? AND student_id = ?"
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

		sm := StudentMgr{}
		qs := Qstu{CardId: card_id}
		r1 := sm.Query(qs, db)
		result, err := stmt.Exec(r[0].ID, r1[0].ID)
		if err != nil {
			fmt.Printf("exec failed! err: %v\n", err)
			return
		}

		id, _ := result.LastInsertId()
		affectRows, _ := result.RowsAffected()
		fmt.Printf("last insert id:%d affect rows:%d\n", id, affectRows)
	}
}
