package models

import (
	"time"
)

type Book struct {
	Id          int64     `db:"id"`
	BookId      string    `db:"book_id"`
	Author      string    `db:"author"`
	BookName    string    `db:"book_name"`
	PublishTime time.Time `db:"publish_time"`
	StockNum    uint      `db:"stock_num"`
}
