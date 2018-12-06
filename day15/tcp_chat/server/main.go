package main

import (
	"day15/tcp_chat/protocol"
	"fmt"
	"io"
	"logging"
	"net"
)

func main() {
	fmt.Println("start server...")

	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		typ, msg, err := protocol.UnPack(conn)
		if err != io.EOF && err != nil {
			logging.Error("read package from failed, close client, err:%v", err)
			return
		}

		err = processMessage(conn, typ, msg)
		if err != nil {
			logging.Error("process message failed, err:%v", err)
			continue
		}

		logging.Debug("process message success, type:%d package:%v", typ, msg)
	}
}

func processMessage(conn net.Conn, typ uint16, msg interface{}) (err error) {
	switch typ {
	case protocol.GET_ROOM_LIST:
		getRoomList, ok := msg.(*protocol.GetRoomList)
		if !ok {
			logging.Error("convert to *protocol.GetRoomList failed, msg:%v", msg)
			return
		}
		return procGetRoomList(conn, typ, getRoomList)
	case protocol.USER_ENTER_ROOM:
		enterRoom, ok := msg.(*protocol.UserEnterRoom)
		if !ok {
			logging.Error("convert to *protocol.UserEnterRoom failed, msg:%v", msg)
			return
		}
		return procEnterRoom(conn, typ, enterRoom)
	case protocol.USER_SEND_TEXT:
		sendText, ok := msg.(*protocol.UserSendText)
		if !ok {
			logging.Error("convert to protocol.GetRoomList failed, message:%#v", msg)
			return
		}
		return procSendText(conn, typ, sendText)
	}
	return
}

func procGetRoomList(conn net.Conn, typ uint16, getRoomList *protocol.GetRoomList) (err error) {
	logging.Debug("start process get room list, user_id:%d", getRoomList.UserId)

	allRoomList := &protocol.AllRoomList{}
	for _, room := range roomMgr.GetRoomList() {
		allRoomList.RoomList = append(allRoomList.RoomList, room)
	}

	data, err := protocol.Pack(protocol.ALL_ROOM_LIST, allRoomList)
	if err != nil {
		logging.Error("pack message failed, data:%#v, err:%v", allRoomList, err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		logging.Error("send data to client failed, err:%v, data:%#v", err, allRoomList)
	}
	return
}

func procEnterRoom(conn net.Conn, typ uint16, enterRoom *protocol.UserEnterRoom) (err error) {
	logging.Debug("start process user enter room, user_id:%d, room_id:%d",
		enterRoom.UserID, enterRoom.RoomID)

	user := NewUser(enterRoom.UserID, enterRoom.UserName, conn)
	roomMgr.EnterRoom(user, enterRoom.RoomID)
	return
}

func procSendText(conn net.Conn, typ uint16, sendText *protocol.UserSendText) (err error) {
	logging.Debug("start process user send text, user_id:%d, room_id:%d, text:%s",
		sendText.UserID, sendText.RoomID, sendText.Content)

	err = roomMgr.SendText(sendText)
	return
}
