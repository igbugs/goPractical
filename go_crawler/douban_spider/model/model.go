package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"logging"
)

var (
	DB *gorm.DB

	username = "gouser"
	password = "123456"
	hostname = "192.168.247.133:3306"
	dbName   = "douban_movice"
)

func init() {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, hostname, dbName))
	if err != nil {
		logging.Error("gorm.Open failed, err: %v", err)
	}

	DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sp_" + defaultTableName
	}
}
