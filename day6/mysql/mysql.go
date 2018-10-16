package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func queryRow(db *sql.DB) {
	id := 1
	row := db.QueryRow("select id, name, age from user where id=?", id)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err == sql.ErrNoRows {
		fmt.Printf("not found data of id %d\n", id)
		return
	}

	if err != nil {
		fmt.Printf("scan row failed! err: %v\n", err)
		return
	}

	fmt.Printf("user: %#v\n", user)
}

func query(db *sql.DB) {
	id := 1
	rows, err := db.Query("select id, name, age from user where id>=?", id)

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err == sql.ErrNoRows {
		fmt.Printf("not found data of id %d\n", id)
		return
	}

	if err != nil {
		fmt.Printf("scan row failed! err: %v\n", err)
		return
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan row failed! err: %v\n", err)
			return
		}
		fmt.Printf("user: %#v\n", user)
	}
}

func insertDb(db *sql.DB) {
	username := "user1000"
	age := 100
	id := 100


	result, err := db.Exec("insert into user(id, name, age) value(?,?,?)", id, username, age)
	if err != nil {
		fmt.Printf("exec failed! err: %v\n", err)
		return
	}

	idnum, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("lastInsert failed! err: %v\n", err)
		return
	}

	affectRow, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affectRow failed! err: %v\n", err)
		return
	}

	fmt.Printf("lastInsert id %d, affectRows: %d\n", idnum, affectRow)
}

func updateDb(db *sql.DB) {
	username := "user1001"
	age := 101
	id := 100

	result, err := db.Exec("update user set name = ?, age = ? where id = ?", username, age, id)
	if err != nil {
		fmt.Printf("exec failed! err: %v\n", err)
		return
	}

	affectRow, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affectRow failed! err: %v\n", err)
		return
	}

	fmt.Printf("affectRows: %d\n", affectRow)
}

func deleteDb(db *sql.DB) {
	id := 100

	result, err := db.Exec("delete FROM user where id = ?", id)
	if err != nil {
		fmt.Printf("exec failed! err: %v\n", err)
		return
	}

	affectRow, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affectRow failed! err: %v\n", err)
		return
	}

	fmt.Printf("affectRows: %d\n", affectRow)
}

func prequery(db *sql.DB) {
	stmt, err := db.Prepare("select id, name, age from user where id>?")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	id := 1
	rows, err := stmt.Query(id)
	defer func() {
		if rows != nil {
			rows.Close()
		}
		if stmt != nil {
			stmt.Close()
		}
	}()

	if err == sql.ErrNoRows {
		fmt.Printf("not found data of id %d\n", id)
		return
	}

	if err != nil {
		fmt.Printf("scan row failed! err: %v\n", err)
		return
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan row failed! err: %v\n", err)
			return
		}
		fmt.Printf("user: %#v\n", user)
	}
}

func prequeryInsert(db *sql.DB) {
	stmt, err := db.Prepare("insert into user(name, age) value(?,?)")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	result, err := stmt.Exec("user1112", 111)

	if err != nil {
		fmt.Printf("exec failed! err: %v\n", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("last insert id failed, err:%v\n", err)
		return
	}

	affectRows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affectRows failed, err:%v\n", err)
		return
	}

	fmt.Printf("last insert id:%d affect rows:%d\n", id, affectRows)
}

func Transaction(Db *sql.DB) {

	tx, err := Db.Begin()
	if err != nil {
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}

	_, err = tx.Exec("insert into user(name, age)values(?, ?)", "user0101", 108)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.Exec("update user set name=?, age=?", "user0101", 108)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}
}

func main() {
	dsn := "gouser:123456@tcp(192.168.247.133:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql failed, err: %v\n", err)
		return
	}

	err = db.Ping()		// 防止 账户密码出错，进行检查
	if err != nil {
		fmt.Printf("ping mysql failed, err: %v\n", err)
		return
	}

	fmt.Printf("connect mysql successfull.\n")

	//queryRow(db)
	query(db)
	//	insertDb(db)
	//	updateDb(db)
	//	deleteDb(db)
	//prequery(db)
	//prequeryInsert(db)
	//Transaction(db)
}
