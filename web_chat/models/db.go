package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"logging"
)

var (
	DB *sqlx.DB
)

func InitDB(dsn string) (err error) {
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		logging.Error("open mysql failed, err:%v", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		logging.Error("ping failed, err:%v", err)
		return
	}

	return
}
