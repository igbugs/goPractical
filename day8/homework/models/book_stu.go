package models

type BookStu struct {
	Id     int64 `db:"id"`
	BookId int64 `db:"book_id"`
	StuId  int64 `db:"student_id"`
}
