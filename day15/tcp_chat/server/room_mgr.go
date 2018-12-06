package main

import (
	"day15/tcp_chat/protocol"
	"logging"
	"sync"
	"sync/atomic"
)

type RoomMgr struct {
	AllRoomList []*RoomInfo
	lock        sync.Mutex
}

var (
	roomMgr *RoomMgr
)

func init() {
	var allRoomList []*protocol.Room
	loveRoom := &protocol.Room{
		RoomID:  1,
		Name:    "谈情说爱",
		RoomCap: 500,
		Desc:    "",
		Online:  0,
	}
	allRoomList = append(allRoomList, loveRoom)

	goRoom := &protocol.Room{
		RoomID:  2,
		Name:    "go 开发论坛",
		RoomCap: 500,
		Desc:    "",
		Online:  0,
	}
	allRoomList = append(allRoomList, goRoom)

	javaRoom := &protocol.Room{
		RoomID:  3,
		Name:    "java 开发论坛",
		RoomCap: 500,
		Desc:    "",
		Online:  0,
	}
	allRoomList = append(allRoomList, javaRoom)

	roomMgr = NewRoomMgr(allRoomList)
}

func NewRoomMgr(roomList []*protocol.Room) (roomMgr *RoomMgr) {
	roomMgr = &RoomMgr{}
	//初始化所有的房间信息
	for _, r := range roomList {
		roomInfo := &RoomInfo{
			room:    r,
			userMap: make(map[uint64]*User, 1024),
		}

		roomMgr.AllRoomList = append(roomMgr.AllRoomList, roomInfo)
	}

	return
}

func (r *RoomMgr) GetRoomList() (roomList []*protocol.Room) {
	for _, roomInfo := range r.AllRoomList {
		roomList = append(roomList, roomInfo.room)
	}

	return
}

func (r *RoomMgr) EnterRoom(user *User, roomId uint64) (err error) {
	var currRoomInfo *RoomInfo
	r.lock.Lock()
	for _, roomInfo := range r.AllRoomList {
		if roomInfo.room.RoomID == roomId {
			currRoomInfo = roomInfo
			break
		}
	}
	r.lock.Unlock()

	if currRoomInfo == nil {
		return user.handleEnterRoomResult(protocol.ERR_INVALID_ROOM_ID, nil)
	}

	atomic.AddUint32(&currRoomInfo.room.Online, 1)
	err = user.handleEnterRoomResult(protocol.SUCCESS, currRoomInfo.room)

	//2. 通知房间里的其他人，有新的用户加入
	broadcastEnterRoom := &protocol.BroadcastEnterRoom{
		EnterUserID:   user.UserID,
		EnterUserName: user.UserName,
		RoomInfo:      currRoomInfo.room,
	}

	msg := &Message{
		Type: protocol.BROADCAST_USER_ENTER_ROOM,
		Body: broadcastEnterRoom,
	}

	for _, otherUser := range currRoomInfo.userMap {
		otherUser.AppendMessage(msg)
	}

	//3. 把当前的用户添加到UserMap里面
	currRoomInfo.userMap[user.UserID] = user
	return
}

func (r *RoomMgr) SendText(sendText *protocol.UserSendText) (err error) {
	var currRoomInfo *RoomInfo
	r.lock.Lock()
	for _, roomInfo := range r.AllRoomList {
		if roomInfo.room.RoomID == sendText.RoomID {
			currRoomInfo = roomInfo
			break
		}
	}
	r.lock.Unlock()

	if currRoomInfo == nil {
		//需要保存一个当前的用户在线的列表
		//return user.handleSendTextResult(protocol.ERR_INVALID_ROOM_ID, nil)
		return
	}

	//2. 从房间里找到这个用户
	user, ok := currRoomInfo.userMap[sendText.UserID]
	if !ok {
		logging.Error("not found user: %v", user)
		return
	}

	//3. 通知当前房间里的其他的用户，由用户发言
	broadcastRecvText := &protocol.UserRecvText{
		Code:           protocol.SUCCESS,
		RoomID:         sendText.RoomID,
		AuthorUserID:   sendText.UserID,
		AuthorUserName: sendText.UserName,
		Content:        sendText.Content,
	}

	msg := &Message{
		Type: protocol.USER_RECV_TEXT,
		Body: broadcastRecvText,
	}

	for _, otherUser := range currRoomInfo.userMap {
		otherUser.AppendMessage(msg)
	}
	return
}
