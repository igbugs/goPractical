package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"database/sql"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func queryRow(db *sqlx.DB) {
	id := 1
	var user User
	err := db.Get(&user,"select id, name, age from user where id=?", id)

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

func query(db *sqlx.DB)  {
	var user []User
	id := 0

	err := db.Select(&user, "select id, name, age from user where id>?", id)
	if err != nil {
		return
	}

	fmt.Printf("user: %#v\n", user)
}

func insertDb(db *sqlx.DB) {
	username := "user1000"
	age := 100

	result, err := db.Exec("insert into user(name, age) value(?,?)", username, age)
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

func main() {
	dsn := "gouser:123456@tcp(192.168.247.133:3306)/test"
	db, err := sqlx.Connect("mysql", dsn)
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
	//query(db)
	insertDb(db)
	//	updateDb(db)
	//	deleteDb(db)
	//prequery(db)
	//prequeryInsert(db)
	//Transaction(db)
}
