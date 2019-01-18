package models

import (
	"logging"
	"time"
)

type Room struct {
	Id         uint64     `db:"id"`
	RoomId     uint64     `db:"room_id"`
	RoomName   string    `db:"room_name"`
	Desc       string    `db:"desc"`
	Online     int       `db:"online"`
	Status     int       `db:"status"`
	Cap        int       `db:"cap"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

func GetAllRoomList() (roomList []*Room, err error) {
	sql := "select id, room_id, room_name, `desc`, online, cap, create_time, update_time from room where status=1"
	err = DB.Select(&roomList, sql)
	if err != nil {
		logging.Error("get all room list failed, err:%v", err)
		return
	}

	return
}

func UpdateRoomOnline(roomId uint64) (err error) {
	sql := "update room set online=online+1 where room_id=?"
	_, err = DB.Exec(sql, roomId)
	if err != nil {
		logging.Error("update room online number failed, err:%v", err)
		return
	}

	return
}