package main

import (
	"day15/tcp_chat/protocol"
	"fmt"
	"logging"
	"net"
)

func getRoomList(conn net.Conn) (roomList *protocol.AllRoomList, err error) {
	var getRoomList = &protocol.GetRoomList{
		UserId: user.UserID,
	}

	data, err := protocol.Pack(protocol.GET_ROOM_LIST, getRoomList)
	if err != nil {
		logging.Error("pack failed, err:%v", err)
		return
	}
	logging.Debug("get room list type, data_len:%d, err:%v", len(data), err)

	_, err = conn.Write(data)
	if err != nil {
		logging.Error("write data failed, err:%v", err)
		return
	}

	typ, result, err := protocol.UnPack(conn)
	if err != nil {
		logging.Error("read message from server failed, message type:%d, err:%v", typ, err)
		return
	}

	if typ != protocol.ALL_ROOM_LIST {
		err = fmt.Errorf("unexpected package, type:%v, data:%v", typ, data)
		logging.Error("unexpected package, type:%v, data:%v", typ, err)
		return
	}

	roomList, ok := result.(*protocol.AllRoomList)
	if !ok {
		err = fmt.Errorf("convert to *protocol.AllRoomList failed")
		logging.Error("err:%v, data:%#v", err, result)
		return
	}

	return
}

func showRoomList(roomList *protocol.AllRoomList) {
	fmt.Printf("===============激情聊天室==================\n")
	fmt.Printf("房间列表\n")
	for _, room := range roomList.RoomList {
		showRoom(room)
	}
	fmt.Printf("please select room id to enter: \n")
}

func showRoom(room *protocol.Room) {
	fmt.Printf("房间编号: %d\n", room.RoomID)
	fmt.Printf("房间名称: %s\n", room.Name)
	fmt.Printf("房间描述: %s\n", room.Desc)
	fmt.Printf("在线人数: %d\n", room.Online)
	fmt.Println()
}

func enterRoom(conn net.Conn, roomId uint64) (roomInfo *protocol.Room, err error) {
	//1. 校验roomid 的合法性
	var validId bool
	for _, room := range roomList.RoomList {
		if room.RoomID == roomId {
			validId = true
			break
		}
	}

	if validId == false {
		err = fmt.Errorf("invalid room id")
		return
	}

	var enterRoom = &protocol.UserEnterRoom{
		RoomID: roomId,
		UserID: user.UserID,
		UserName: user.UserName,
	}

	//打包成网络字节流
	data, err := protocol.Pack(protocol.USER_ENTER_ROOM, enterRoom)
	if err != nil {
		logging.Error("pack failed, err:%v", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		logging.Error("write data failed, err:%v", err)
		return
	}

	typ, result, err := protocol.UnPack(conn)
	if err != nil {
		logging.Error("read message from server failed, message type:%d err:%v", typ, err)
		return
	}

	if typ != protocol.USER_ENTER_ROOM_RESP {
		err = fmt.Errorf("unexpected package, type:%v, data:%v", typ, data)
		logging.Error("unexpected package, type:%v, data:%v", typ, err)
		return
	}

	resp, ok := result.(*protocol.UserEnterRoomResp)
	if !ok {
		err = fmt.Errorf("convert to *protocol.UserEnterRoomResp failed")
		logging.Error("err:%v, data:%#v", err, result)
		return
	}

	if resp.Code != protocol.SUCCESS {
		err = fmt.Errorf("enter room failed, code:%d", resp.Code)
		logging.Error("enter room failed, code:%d", resp.Code)
		return
	}

	fmt.Printf("enter room success\n")
	showRoom(resp.RoomInfo)
	roomInfo = resp.RoomInfo
	return
}
