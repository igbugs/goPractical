package models

import (
	"logging"
	"time"
)

type User struct {
	Id         uint64    `db:"id"`
	UserId     uint64    `db:"user_id"`
	Username   string    `db:"username"`
	Nickname   string    `db:"nickname"`
	Sex        int       `db:"sex"`
	Password   string    `db:"password"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

func Register(user *User) (err error) {
	sql := "insert into user (user_id, username, nickname, sex, password) values (?, ?, ?, ?, ?)"
	_, err = DB.Exec(sql, user.UserId, user.Username, user.Nickname, user.Sex, user.Password)
	if err != nil {
		logging.Error("insert failed, err:%v", err)
		return
	}
	return
}

func GetUserByName(username string) (user *User, err error) {
	user = &User{}
	sql := "select user_id, username, nickname, sex, password from user where username=?"
	err = DB.Get(user, sql, username)
	if err != nil {
		logging.Error("get user info by username:%s failed, err:%v", username, err)
		return
	}

	return
}

func GetUserById(id int64) (user *User, err error) {
	user = &User{}
	sql := "select user_id, username, nickname, sex, password from user where username=?"
	err = DB.Get(user, sql, id)
	if err != nil {
		logging.Error("get user info by username:%s failed, err:%v", id, err)
		return
	}

	return
}
