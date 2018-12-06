package main

import (
	"day15/tcp_chat/protocol"
	"logging"
	"net"
)

type Message struct {
	Type uint16
	Body interface{}
}

type User struct {
	UserID   uint64
	UserName string
	Conn     net.Conn
	OutBox   chan *Message
}

func NewUser(userId uint64, username string, conn net.Conn) (user *User) {
	user = &User{
		UserID: userId,
		UserName: username,
		Conn: conn,
		OutBox: make(chan *Message, 1024),
	}

	go user.sendMessage()
	return
}

func (u *User) sendMessage()  {
	for msg := range u.OutBox {
		data, err := protocol.Pack(msg.Type, msg.Body)
		if err != nil {
			logging.Error("pack message:%#v, err:%v", msg, err)
			continue
		}

		_, err = u.Conn.Write(data)
		if err != nil {
			logging.Error("send to client failed, err:%v", err)
			continue
		}
	}
}

func (u *User) handleEnterRoomResult(code int, room *protocol.Room) (err error) {
	enterRoomResp := &protocol.UserEnterRoomResp{
		Code: code,
		RoomInfo: room,
	}

	data, err := protocol.Pack(protocol.USER_ENTER_ROOM_RESP, enterRoomResp)
	if err != nil {
		logging.Error("pack message protocol.USER_ENTER_ROOM_RESP failed, err:%v", err)
		return
	}

	_, err = u.Conn.Write(data)
	return
}

func (u *User) handleSendTextResult(code int, room *protocol.Room) (err error) {
	recvText := &protocol.UserRecvText{
		Code: code,
	}

	data, err := protocol.Pack(protocol.USER_RECV_TEXT, recvText)
	if err != nil {
		logging.Error("pack message protocol.USER_RECV_TEXT failed, err:%v", err)
		return
	}

	_, err = u.Conn.Write(data)
	return
}

func (u *User) AppendMessage(msg *Message)  {
	select {
	case u.OutBox <- msg:
	default:
		logging.Error("user message chan fill")
		return
	}
}
