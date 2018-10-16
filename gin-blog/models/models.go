package models

import (
	"github.com/jinzhu/gorm"
	"gin-blog/pkg/setting"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gin-blog/pkg/logging"
)

var db *gorm.DB

type Model struct {
	ID int			`gorm:"primary_key" json:"id"`
	CreatedOn int	`json:"created_on"`
	ModifiedOn int	`json:"modified_on"`
}

func init()  {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		logging.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("type").String()
	dbName = sec.Key("name").String()
	user = sec.Key("user").String()
	password = sec.Key("password").String()
	host = sec.Key("host").String()
	tablePrefix = sec.Key("table_prefix").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName))
	if err != nil {
		logging.Info(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB()  {
	defer db.Close()
}
