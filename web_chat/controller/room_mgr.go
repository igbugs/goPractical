package controller

import (
	"fmt"
	"logging"
	"sort"
	"time"
	"web_chat/models"
)

var (
	roomMgr *RoomMgr
)

type RoomList []*models.Room

func (rl RoomList) Len() int {
	roomList := []*models.Room(rl)
	return len(roomList)
}

func (rl RoomList) Less(i, j int) bool {
	roomList := []*models.Room(rl)
	r1 := roomList[i]
	r2 := roomList[j]

	return r1.Online > r2.Online
}

func (rl RoomList) Swap(i, j int) {
	roomList := []*models.Room(rl)
	roomList[i], roomList[j] = roomList[j], roomList[i]
}

type RoomInfo struct {
	Room *models.Room
	// 以user.User.UserId 为key值
	UserMap map[uint64]*UserInfo
}

func (r *RoomInfo) DeleteUser(user *UserInfo) {
	delete(r.UserMap, user.User.UserId)
}

func (r *RoomInfo) AddUser(user *UserInfo) (alreadyLogin bool) {
	_, alreadyLogin = r.UserMap[user.User.UserId]
	if alreadyLogin {
		return
	}

	r.UserMap[user.User.UserId] = user
	err := models.UpdateRoomOnline(r.Room.RoomId)
	if err != nil {
		logging.Error("update room online failed, room_id:%d err:%v", r.Room.RoomId, err)
		return
	}

	return
}

type RoomMgr struct {
	RoomMap map[uint64]*RoomInfo
}

func NewRoomMgr() *RoomMgr  {
	return &RoomMgr{
		RoomMap: make(map[uint64]*RoomInfo, 16),
	}
}

func (r *RoomMgr) Init(roomList []*models.Room) (err error) {
	for _, room := range roomList {
		roomInfo := &RoomInfo{
			Room: room,
			UserMap: make(map[uint64]*UserInfo, 1024),
		}

		r.RoomMap[room.RoomId] = roomInfo
	}
	return
}

func (r *RoomMgr) GetRoom(roomId uint64) (roomInfo *RoomInfo, err error) {
	roomInfo, ok := r.RoomMap[roomId]
	if !ok {
		logging.Error("room not exists, room_id:%d", roomId)
		err = fmt.Errorf("room not exists, room_id:%d", roomId)
		return
	}

	return
}

func (r *RoomMgr) GetRoomList() (roomList []*models.Room, err error){
	for _, v := range r.RoomMap {
		roomList = append(roomList, v.Room)
	}

	var sortRoomList = RoomList(roomList)
	sort.Sort(sortRoomList)
	for _, r := range sortRoomList {
		logging.Debug("sort result:%#v", r)
	}
	roomList = []*models.Room(sortRoomList)
	return
}

func (r *RoomMgr) SyncRoomList() {
	for {
		time.Sleep(time.Second)
		roomList, err := models.GetAllRoomList()
		if err != nil {
			logging.Error("get all room list failed, err:%v", err)
			continue
		}

		for _, room := range roomList {
			_, ok := r.RoomMap[room.RoomId]
			if !ok {
				roomInfo := &RoomInfo{
					Room: room,
					UserMap: make(map[uint64]*UserInfo, 1024),
				}

				r.RoomMap[room.RoomId] = roomInfo
				continue
			}
			r.RoomMap[room.RoomId].Room = room
		}
	}
}

func InitRoomMgr() (err error) {
	roomMgr = NewRoomMgr()
	roomList, err := models.GetAllRoomList()
	if err != nil {
		return
	}

	err = roomMgr.Init(roomList)
	go roomMgr.SyncRoomList()
	return
}

