package models

type Student struct {
	Id int64 `db:"id"`
	StudentId  string `db:"student_id"`
	PassWord string	`db:"password"`
	Name string `db:"stu_name"`
	Age string 	`db:"age"`
	Sex string	`db:"sex"`
	Grade string `db:"grade"`
}
