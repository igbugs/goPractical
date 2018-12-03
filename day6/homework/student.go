package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type Student struct {
	ID     string `db:"id"`
	CardId string `db:"card_id"`
	Name   string `db:"name"`
	Sex    string `db:"sex"`
	Grade  string `db:"grade"`
}

type StudentMgr struct{}

type Qstu struct {
	ID     string
	CardId string
	Name   string
	Sex    string
	Grade  string
}

func (bm StudentMgr) Query(qs Qstu, db *sql.DB) []Student {
	sqlStr := "SELECT * FROM student WHERE 1=1"

	if qs.ID != "" {
		sqlStr += " AND id=" + qs.ID
	}
	if qs.CardId != "" {
		sqlStr += " AND card_id='" + qs.CardId + "'"
	}
	if qs.Name != "" {
		sqlStr += " AND name='" + qs.Name + "'"
	}
	if qs.Sex != "" {
		sqlStr += " AND number='" + qs.Sex + "'"
	}
	if qs.Grade != "" {
		sqlStr += " AND author='" + qs.Grade + "'"
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

	var ss []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.CardId, &student.Name, &student.Sex, &student.Grade)
		if err != nil {
			fmt.Printf("scan row failed! err: %v\n", err)
		}
		ss = append(ss, student)
	}
	return ss
}

func (bm StudentMgr) AddStudent(s Student, db *sql.DB) {
	var qs Qstu
	if s.CardId != "" {
		qs.CardId = s.CardId
	}

	r := bm.Query(qs, db)
	if len(r) == 0 {
		sqlStr := "INSERT INTO student (card_id, name, sex, grade) VALUE (?, ?, ?, ?)"
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

		result, err := stmt.Exec(s.CardId, s.Name, s.Sex, s.Grade)
		if err != nil {
			fmt.Printf("exec failed! err: %v\n", err)
			return
		}

		id, _ := result.LastInsertId()
		affectRows, _ := result.RowsAffected()
		fmt.Printf("last insert id:%d affect rows:%d\n", id, affectRows)
	} else {
		fmt.Printf("student card_id: %s is exist!", s.CardId)
		return
	}
}

func (bm StudentMgr) DelStudent(s Student, db *sql.DB) {
	var qs Qstu
	if s.CardId != "" {
		qs.CardId = s.CardId
	}

	r := bm.Query(qs, db)
	if len(r) != 0 {
		sqlStr := "DELETE FROM student WHERE card_id=?"
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

		result, err := stmt.Exec(s.CardId)
		if err != nil {
			fmt.Printf("exec failed! err: %v\n", err)
			return
		}

		affectRows, _ := result.RowsAffected()
		fmt.Printf("affect rows:%d\n", affectRows)
	} else {
		fmt.Printf("student card_id: %s is exist!", s.CardId)
		return
	}
}

func (bm StudentMgr) UpdateStudent(s Student, db *sql.DB) {
	var qs Qstu
	if s.CardId != "" {
		qs.CardId = s.CardId
	}
	r := bm.Query(qs, db)
	if len(r) != 0 {
		sqlStr := "UPDATE student SET"
		rs := r[0]
		if rs.Name != s.Name && s.Name != "" {
			sqlStr += " name='" + s.Name + "',"
		}
		if rs.Sex != s.Sex && s.Sex != "" {
			sqlStr += " sex='" + s.Sex + "',"
		}
		if rs.Grade != s.Grade && s.Grade != "" {
			sqlStr += " grade='" + s.Grade + "',"
		}

		sqlStr = strings.Trim(sqlStr, ",") + " WHERE card_id='" + s.CardId + "'"
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
		fmt.Printf("student card_id: %s is exist!", s.CardId)
		return
	}
}
